package simulation

import (
	stdError "errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aimiddleware"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aimiddleware/reverts"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aiop"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/errors"
)

// SimulateValidation makes a static call to Aimiddleware.simulateValidation(aiop) and returns the
// results without any state changes.
func SimulateValidation(
	rpc *rpc.Client,
	aiMiddleware common.Address,
	op *aiop.AiOperation,
) (*reverts.ValidationResultRevert, error) {
	ep, err := aimiddleware.NewAimiddleware(aiMiddleware, ethclient.NewClient(rpc))
	if err != nil {
		return nil, err
	}

	var res []interface{}
	rawCaller := &aimiddleware.AimiddlewareRaw{Contract: ep}
	err = rawCaller.Call(nil, &res, "simulateValidation", aimiddleware.AiOperation(*op))
	if err == nil {
		return nil, stdError.New("unexpected result from simulateValidation")
	}

	sim, simErr := reverts.NewValidationResult(err)
	if simErr != nil {
		fo, foErr := reverts.NewFailedOp(err)
		if foErr != nil {
			return nil, fmt.Errorf("%s, %s", simErr, foErr)
		}
		return nil, errors.NewRPCError(errors.REJECTED_BY_EP_OR_ACCOUNT, fo.Reason, fo)
	}

	return sim, nil
}
