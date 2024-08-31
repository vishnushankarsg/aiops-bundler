package methods

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aiop"
)

var (
	HandleOpsMethod = abi.NewMethod(
		"handleOps",
		"handleOps",
		abi.Function,
		"",
		false,
		false,
		abi.Arguments{
			{Name: "ops", Type: aiop.AiOpArr},
			{Name: "beneficiary", Type: address},
		},
		nil,
	)
	HandleOpsSelector = hexutil.Encode(HandleOpsMethod.ID)
)
