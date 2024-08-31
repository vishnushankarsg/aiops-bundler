package simulation

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aimiddleware/methods"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aiop"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/tracer"
)

type knownEntity map[string]struct {
	Address  common.Address
	Info     tracer.CallFromAiMiddlewareInfo
	IsStaked bool
}

func newKnownEntity(
	op *aiop.AiOperation,
	res *tracer.BundlerCollectorReturn,
	stakes EntityStakes,
) (knownEntity, error) {
	si := tracer.CallFromAiMiddlewareInfo{}
	fi := tracer.CallFromAiMiddlewareInfo{}
	pi := tracer.CallFromAiMiddlewareInfo{}
	for _, c := range res.CallsFromAiMiddleware {
		switch c.TopLevelTargetAddress {
		case op.Sender:
			si = c
		case op.GetPaymaster():
			pi = c
		default:
			if c.TopLevelMethodSig.String() == methods.CreateSenderSelector {
				fi = c
			}
		}
	}

	return knownEntity{
		"account": {
			Address:  op.Sender,
			Info:     si,
			IsStaked: stakes[op.Sender] != nil && stakes[op.Sender].Staked,
		},
		"factory": {
			Address:  op.GetFactory(),
			Info:     fi,
			IsStaked: stakes[op.GetFactory()] != nil && stakes[op.GetFactory()].Staked,
		},
		"paymaster": {
			Address:  op.GetPaymaster(),
			Info:     pi,
			IsStaked: stakes[op.GetPaymaster()] != nil && stakes[op.GetPaymaster()].Staked,
		},
	}, nil
}

func addr2KnownEntity(op *aiop.AiOperation, addr common.Address) string {
	if addr == op.GetFactory() {
		return "factory"
	} else if addr == op.Sender {
		return "account"
	} else if addr == op.GetPaymaster() {
		return "paymaster"
	} else {
		return addr.String()
	}
}
