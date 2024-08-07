package execution

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/reverts"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/utils"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/errors"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/state"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/tracer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethRpc "github.com/ethereum/go-ethereum/rpc"
)

type TraceInput struct {
	Rpc          *ethRpc.Client
	AiMiddleware common.Address
	Op           *aiop.AiOperation
	Sos          state.OverrideSet
	ChainID      *big.Int
	Tracer       string

	// Optional params for simulateHandleOps
	Target      common.Address
	Data        []byte
	TraceFeeCap *big.Int
}

type TraceOutput struct {
	Trace  *tracer.BundlerExecutionReturn
	Result *reverts.ExecutionResultRevert
	Event  *aimiddleware.AimiddlewareAiOperationEvent
}

func parseAiOperationEvent(
	aiMiddleware common.Address,
	ep *aimiddleware.Aimiddleware,
	log *tracer.LogInfo,
) (*aimiddleware.AimiddlewareAiOperationEvent, error) {
	if log == nil {
		return nil, nil
	}

	topics := []common.Hash{}
	for _, topic := range log.Topics {
		topics = append(topics, common.HexToHash(topic))
	}
	data, err := hexutil.Decode(log.Data)
	if err != nil {
		return nil, err
	}

	ev, err := ep.ParseAiOperationEvent(types.Log{
		Address: aiMiddleware,
		Topics:  topics,
		Data:    data,
	})
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func TraceSimulateHandleOp(in *TraceInput) (*TraceOutput, error) {
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
	mf := in.Op.MaxFeePerGas
	if in.TraceFeeCap != nil {
		mf = in.TraceFeeCap
	}
	tx, err := ep.SimulateHandleOp(auth, aimiddleware.AiOperation(*in.Op), in.Target, in.Data)
	if err != nil {
		return nil, err
	}
	t := tracer.Loaded.BundlerExecutionTracer
	if in.Tracer != "" {
		t = in.Tracer
	}
	out := &TraceOutput{}

	var res tracer.BundlerExecutionReturn
	req := utils.TraceCallReq{
		From:         common.HexToAddress("0x"),
		To:           in.AiMiddleware,
		Data:         tx.Data(),
		MaxFeePerGas: hexutil.Big(*mf),
	}
	opts := utils.TraceCallOpts{
		Tracer:         t,
		StateOverrides: state.WithMaxBalanceOverride(common.HexToAddress("0x"), in.Sos),
	}
	if err := in.Rpc.CallContext(context.Background(), &res, "debug_traceCall", &req, "latest", &opts); err != nil {
		return nil, err
	}
	outErr, err := errors.ParseHexToRpcDataError(res.Output)
	if err != nil {
		return nil, err
	}
	if res.ValidationOOG {
		return nil, errors.NewRPCError(errors.EXECUTION_REVERTED, "validation OOG", nil)
	}
	out.Trace = &res

	sim, simErr := reverts.NewExecutionResult(outErr)
	if simErr != nil {
		fo, foErr := reverts.NewFailedOp(outErr)
		if foErr != nil && res.Error != "" {
			return nil, errors.NewRPCError(errors.EXECUTION_REVERTED, res.Error, nil)
		} else if foErr != nil {
			return nil, fmt.Errorf("%s, %s", simErr, foErr)
		}
		return nil, errors.NewRPCError(errors.REJECTED_BY_EP_OR_ACCOUNT, fo.Reason, fo)
	}
	out.Result = sim

	ev, err := parseAiOperationEvent(in.AiMiddleware, ep, res.AiOperationEvent)
	if err != nil {
		return out, err
	}
	out.Event = ev

	if !ev.Success && len(res.Reverts) != 0 {
		data, err := hexutil.Decode(res.Reverts[len(res.Reverts)-1])
		if err != nil {
			return out, err
		}

		if len(data) == 0 {
			if res.ExecutionOOG {
				return out, errors.NewRPCError(errors.EXECUTION_REVERTED, "execution OOG", nil)
			}
			return out, errors.NewRPCError(errors.EXECUTION_REVERTED, "execution reverted", nil)
		}

		reason, revErr := errors.DecodeRevert(data)
		if revErr != nil {
			code, panErr := errors.DecodePanic(data)
			if panErr != nil {
				return nil, errors.NewRPCError(
					errors.EXECUTION_REVERTED,
					"execution reverted with data",
					hexutil.Encode(data),
				)
			}

			return out, errors.NewRPCError(
				errors.EXECUTION_REVERTED,
				fmt.Sprintf("panic encountered: %s", code),
				code,
			)
		}
		return out, errors.NewRPCError(errors.EXECUTION_REVERTED, reason, reason)
	}

	return out, nil
}
