package methods

import (
	"errors"
	"fmt"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ValidatePaymasterAiOpMethod = abi.NewMethod(
		"validatePaymasterAiOp",
		"validatePaymasterAiOp",
		abi.Function,
		"",
		false,
		false,
		abi.Arguments{
			{Name: "aiOp", Type: aiop.AiOpType},
			{Name: "aiOpHash", Type: bytes32},
			{Name: "maxCost", Type: uint256},
		},
		abi.Arguments{
			{Name: "context", Type: bytes},
			{Name: "deadline", Type: uint256},
		},
	)
	ValidatePaymasterAiOpSelector = hexutil.Encode(ValidatePaymasterAiOpMethod.ID)
)

type validatePaymasterAiOpOutput struct {
	Context []byte
}

func DecodeValidatePaymasterAiOpOutput(ret any) (*validatePaymasterAiOpOutput, error) {
	hex, ok := ret.(string)
	if !ok {
		return nil, errors.New("validatePaymasterAiOp: cannot assert type: hex is not of type string")
	}
	data, err := hexutil.Decode(hex)
	if err != nil {
		return nil, fmt.Errorf("validatePaymasterAiOp: %s", err)
	}

	args, err := ValidatePaymasterAiOpMethod.Outputs.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("validatePaymasterAiOp: %s", err)
	}
	if len(args) != 2 {
		return nil, fmt.Errorf(
			"validatePaymasterAiOp: invalid args length: expected 2, got %d",
			len(args),
		)
	}

	ctx, ok := args[0].([]byte)
	if !ok {
		return nil, errors.New("validatePaymasterAiOp: cannot assert type: hex is not of type string")
	}

	return &validatePaymasterAiOpOutput{
		Context: ctx,
	}, nil
}
