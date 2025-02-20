package simulation

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware/methods"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware/utils"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/altmempools"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/state"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/tracer"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type TraceInput struct {
	Rpc                *rpc.Client
	AiMiddleware       common.Address
	Op                 *aiop.AiOperation
	ChainID            *big.Int
	IsRIP7212Supported bool
	Tracer             string
	Stakes             EntityStakes
	AltMempools        *altmempools.Directory
}

type TraceOutput struct {
	TouchedContracts []common.Address
	AltMempoolIds    []string
}

// TraceSimulateValidation makes a debug_traceCall to Aimiddleware.simulateValidation(aiop) and returns
// information related to the validation phase of a AiOperation.
func TraceSimulateValidation(in *TraceInput) (*TraceOutput, error) {
	ep, err := aimiddleware.NewAimiddleware(in.AiMiddleware, ethclient.NewClient(in.Rpc))
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(utils.DummyPk, in.ChainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = math.MaxUint64
	auth.NoSend = true
	tx, err := ep.SimulateValidation(auth, aimiddleware.AiOperation(*in.Op))
	if err != nil {
		return nil, err
	}
	t := tracer.Loaded.BundlerCollectorTracer
	if in.Tracer != "" {
		t = in.Tracer
	}

	var res tracer.BundlerCollectorReturn
	req := utils.TraceCallReq{
		From:         common.HexToAddress("0x"),
		To:           in.AiMiddleware,
		Data:         tx.Data(),
		MaxFeePerGas: hexutil.Big(*in.Op.MaxFeePerGas),
	}
	opts := utils.TraceCallOpts{
		Tracer:         t,
		StateOverrides: state.WithMaxBalanceOverride(common.HexToAddress("0x"), nil),
	}
	if err := in.Rpc.CallContext(context.Background(), &res, "debug_traceCall", &req, "latest", &opts); err != nil {
		return nil, err
	}

	knownEntity, err := newKnownEntity(in.Op, &res, in.Stakes)
	altMempoolIds := []string{}
	if err != nil {
		return nil, err
	}

	ic := mapset.NewSet[common.Address]()
	for title, entity := range knownEntity {
		if entity.Info.OOG {
			return nil, fmt.Errorf("%s OOG", title)
		}
		if _, ok := entity.Info.ExtCodeAccessInfo[in.AiMiddleware]; ok {
			return nil, fmt.Errorf("%s has forbidden EXTCODE* access to the AiMiddleware", title)
		}
		for opcode := range entity.Info.Opcodes {
			if bannedOpCodes.Contains(opcode) {
				return nil, fmt.Errorf("%s uses banned opcode: %s", title, opcode)
			}

			if !entity.IsStaked && bannedUnstakedOpCodes.Contains(opcode) {
				return nil, fmt.Errorf("unstaked %s uses banned opcode: %s", title, opcode)
			}
		}

		ic.Add(entity.Address)
		for addr := range entity.Info.ContractSize {
			ic.Add(addr)
		}
	}

	create2Count, ok := knownEntity["factory"].Info.Opcodes[create2OpCode]
	if ok && (create2Count > 1 || len(in.Op.InitCode) == 0) {
		return nil, fmt.Errorf("factory with too many %s", create2OpCode)
	}
	_, ok = knownEntity["account"].Info.Opcodes[create2OpCode]
	if ok {
		return nil, fmt.Errorf("account uses banned opcode: %s", create2OpCode)
	}
	_, ok = knownEntity["paymaster"].Info.Opcodes[create2OpCode]
	if ok {
		return nil, fmt.Errorf("paymaster uses banned opcode: %s", create2OpCode)
	}

	slotsByEntity := newStorageSlotsByEntity(in.Stakes, res.Keccak)
	for title, entity := range knownEntity {
		v := &storageSlotsValidator{
			Op:                    in.Op,
			AiMiddleware:          in.AiMiddleware,
			IsRIP7212Supported:    in.IsRIP7212Supported,
			AltMempools:           in.AltMempools,
			SenderSlots:           slotsByEntity[in.Op.Sender],
			FactoryIsStaked:       knownEntity["factory"].IsStaked,
			EntityName:            title,
			EntityAddr:            entity.Address,
			EntityAccessMap:       entity.Info.Access,
			EntityContractSizeMap: entity.Info.ContractSize,
			EntitySlots:           slotsByEntity[entity.Address],
			EntityIsStaked:        entity.IsStaked,
		}
		if ids, err := v.Process(); err != nil {
			return nil, err
		} else {
			altMempoolIds = append(altMempoolIds, ids...)
		}
	}

	callStack := newCallStack(res.Calls)
	for _, call := range callStack {
		if call.Method == methods.ValidatePaymasterAiOpSelector {
			out, err := methods.DecodeValidatePaymasterAiOpOutput(call.Return)
			if err != nil {
				return nil, fmt.Errorf(
					"unexpected tracing result for op: %s, %s",
					in.Op.GetAiOpHash(in.AiMiddleware, in.ChainID),
					err,
				)
			}

			if len(out.Context) != 0 && !knownEntity["paymaster"].IsStaked {
				return nil, errors.New("unstaked paymaster must not return context")
			}
		} else if call.To == in.AiMiddleware && call.Method == methods.BalanceOfSelector {
			return nil, fmt.Errorf(
				"%s cannot call balanceOf on AiMiddleware",
				addr2KnownEntity(in.Op, call.From),
			)
		} else if call.To != in.AiMiddleware && call.Value.Cmp(common.Big0) == 1 {
			return nil, fmt.Errorf(
				"%s has a forbidden value transfer to %s",
				addr2KnownEntity(in.Op, call.From),
				addr2KnownEntity(in.Op, call.To),
			)
		}
	}

	return &TraceOutput{
		TouchedContracts: ic.ToSlice(),
		AltMempoolIds:    altMempoolIds,
	}, nil
}
