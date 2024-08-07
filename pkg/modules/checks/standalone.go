// Package checks implements modules for running an array of standard validations for both the Client and
// Bundler.
package checks

import (
	"math/big"
	"time"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/simulation"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/altmempools"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/errors"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules/entities"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules/gasprice"
	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
)

// Standalone exposes modules to perform basic Client and Bundler checks as specified in EIP-4337. It is
// intended for bundlers that are independent of an Ethereum node and hence relies on a given ethClient to
// query blockchain state.
type Standalone struct {
	db                 *badger.DB
	rpc                *rpc.Client
	eth                *ethclient.Client
	ov                 *gas.Overhead
	alt                *altmempools.Directory
	maxVerificationGas *big.Int
	maxBatchGasLimit   *big.Int
	isRIP7212Supported bool
	tracer             string
	repConst           *entities.ReputationConstants
}

// New returns a Standalone instance with methods that can be used in Client and Bundler modules to perform
// standard checks as specified in EIP-4337.
func New(
	db *badger.DB,
	rpc *rpc.Client,
	ov *gas.Overhead,
	alt *altmempools.Directory,
	maxVerificationGas *big.Int,
	maxBatchGasLimit *big.Int,
	isRIP7212Supported bool,
	tracer string,
	repConst *entities.ReputationConstants,
) *Standalone {
	eth := ethclient.NewClient(rpc)
	return &Standalone{
		db,
		rpc,
		eth,
		ov,
		alt,
		maxVerificationGas,
		maxBatchGasLimit,
		isRIP7212Supported,
		tracer,
		repConst,
	}
}

// ValidateOpValues returns a AiOpHandler that runs through some first line sanity checks for new AiOps
// received by the Client. This should be one of the first modules executed by the Client.
func (s *Standalone) ValidateOpValues() modules.AiOpHandlerFunc {
	return func(ctx *modules.AiOpHandlerCtx) error {
		gc := getCodeWithEthClient(s.eth)

		g := new(errgroup.Group)
		g.Go(func() error { return ValidateSender(ctx.AiOp, gc) })
		g.Go(func() error { return ValidateInitCode(ctx.AiOp) })
		g.Go(func() error { return ValidateVerificationGas(ctx.AiOp, s.ov, s.maxVerificationGas) })
		g.Go(func() error { return ValidatePaymasterAndData(ctx.AiOp, ctx.GetPaymasterDepositInfo(), gc) })
		g.Go(func() error { return ValidateCallGasLimit(ctx.AiOp, s.ov) })
		g.Go(func() error { return ValidateFeePerGas(ctx.AiOp, gasprice.GetBaseFeeWithEthClient(s.eth)) })
		g.Go(func() error { return ValidatePendingOps(ctx.AiOp, ctx.GetPendingSenderOps()) })
		g.Go(func() error { return ValidateGasAvailable(ctx.AiOp, s.maxBatchGasLimit) })

		if err := g.Wait(); err != nil {
			return errors.NewRPCError(errors.INVALID_FIELDS, err.Error(), err.Error())
		}
		return nil
	}
}

// SimulateOp returns a AiOpHandler that runs through simulation of new AiOps with the AiMiddleware.
func (s *Standalone) SimulateOp() modules.AiOpHandlerFunc {
	return func(ctx *modules.AiOpHandlerCtx) error {
		gc := getCodeWithEthClient(s.eth)
		g := new(errgroup.Group)
		g.Go(func() error {
			sim, err := simulation.SimulateValidation(s.rpc, ctx.AiMiddleware, ctx.AiOp)

			if err != nil {
				return errors.NewRPCError(errors.REJECTED_BY_EP_OR_ACCOUNT, err.Error(), err.Error())
			}
			if sim.ReturnInfo.SigFailed {
				return errors.NewRPCError(
					errors.INVALID_SIGNATURE,
					"Invalid AiOp signature or paymaster signature",
					nil,
				)
			}
			if sim.ReturnInfo.ValidUntil.Cmp(common.Big0) != 0 &&
				time.Now().Unix() >= sim.ReturnInfo.ValidUntil.Int64()-30 {
				return errors.NewRPCError(
					errors.SHORT_DEADLINE,
					"expires too soon",
					nil,
				)
			}
			return nil
		})
		g.Go(func() error {
			out, err := simulation.TraceSimulateValidation(&simulation.TraceInput{
				Rpc:                s.rpc,
				AiMiddleware:       ctx.AiMiddleware,
				AltMempools:        s.alt,
				Op:                 ctx.AiOp,
				ChainID:            ctx.ChainID,
				IsRIP7212Supported: s.isRIP7212Supported,
				Tracer:             s.tracer,
				Stakes: simulation.EntityStakes{
					ctx.AiOp.Sender:         ctx.GetSenderDepositInfo(),
					ctx.AiOp.GetFactory():   ctx.GetFactoryDepositInfo(),
					ctx.AiOp.GetPaymaster(): ctx.GetPaymasterDepositInfo(),
				},
			})
			if err != nil {
				return errors.NewRPCError(errors.BANNED_OPCODE, err.Error(), err.Error())
			}

			ch, err := getCodeHashes(out.TouchedContracts, gc)
			if err != nil {
				return errors.NewRPCError(errors.BANNED_OPCODE, err.Error(), err.Error())
			}
			return saveCodeHashes(s.db, ctx.AiOp.GetAiOpHash(ctx.AiMiddleware, ctx.ChainID), ch)
		})

		return g.Wait()
	}
}

// CodeHashes returns a BatchHandler that verifies the code for any interacted contracts has not changed since
// the first simulation.
func (s *Standalone) CodeHashes() modules.BatchHandlerFunc {
	return func(ctx *modules.BatchHandlerCtx) error {
		gc := getCodeWithEthClient(s.eth)

		end := len(ctx.Batch) - 1
		for i := end; i >= 0; i-- {
			op := ctx.Batch[i]
			chs, err := getSavedCodeHashes(s.db, op.GetAiOpHash(ctx.AiMiddleware, ctx.ChainID))
			if err != nil {
				return err
			}

			changed, err := hasCodeHashChanges(chs, gc)
			if err != nil {
				return err
			}
			if changed {
				ctx.MarkOpIndexForRemoval(i, "code hash changed")
			}
		}
		return nil
	}
}

// PaymasterDeposit returns a BatchHandler that tracks each paymaster in the batch and ensures it has enough
// deposit to pay for all the AiOps that use it.
func (s *Standalone) PaymasterDeposit() modules.BatchHandlerFunc {
	return func(ctx *modules.BatchHandlerCtx) error {
		ep, err := aimiddleware.NewAimiddleware(ctx.AiMiddleware, s.eth)
		if err != nil {
			return err
		}

		deps := make(map[common.Address]*big.Int)
		for i, op := range ctx.Batch {
			pm := op.GetPaymaster()
			if pm == common.HexToAddress("0x") {
				continue
			}

			if _, ok := deps[pm]; !ok {
				dep, err := ep.GetDepositInfo(nil, pm)
				if err != nil {
					return err
				}

				deps[pm] = dep.Deposit
			}

			deps[pm] = big.NewInt(0).Sub(deps[pm], op.GetMaxPrefund())
			if deps[pm].Cmp(common.Big0) < 0 {
				ctx.MarkOpIndexForRemoval(i, "insufficient paymaster deposit")
			}
		}

		return nil
	}
}

// TODO: Implement
func (s *Standalone) SimulateBatch() modules.BatchHandlerFunc {
	return func(ctx *modules.BatchHandlerCtx) error {
		return nil
	}
}

// Clean returns a BatchHandler that clears the DB of data that is no longer required. This should be one of
// the last modules executed by the Bundler.
func (s *Standalone) Clean() modules.BatchHandlerFunc {
	return func(ctx *modules.BatchHandlerCtx) error {
		all := append([]*aiop.AiOperation{}, ctx.Batch...)
		for _, item := range ctx.PendingRemoval {
			all = append(all, item.Op)
		}
		hashes := []common.Hash{}
		for _, op := range all {
			hashes = append(hashes, op.GetAiOpHash(ctx.AiMiddleware, ctx.ChainID))
		}

		return removeSavedCodeHashes(s.db, hashes...)
	}
}
