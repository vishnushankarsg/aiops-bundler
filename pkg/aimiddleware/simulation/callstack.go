package simulation

import (
	"math/big"
	"strings"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/utils"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/tracer"
	"github.com/ethereum/go-ethereum/common"
)

type callEntry struct {
	From   common.Address
	To     common.Address
	Value  *big.Int
	Type   string
	Method string
	Revert any
	Return any
}

func newCallStack(calls []tracer.CallInfo) []*callEntry {
	out := []*callEntry{}
	stack := utils.NewStack[tracer.CallInfo]()
	for _, call := range calls {
		if call.Type == revertOpCode || call.Type == returnOpCode {
			top, _ := stack.Pop()

			if strings.Contains(top.Type, "CREATE") {
				// TODO: implement
			} else if call.Type == revertOpCode {
				// TODO: implement
			} else {
				v, ok := big.NewInt(0).SetString(top.Value, 0)
				if !ok {
					v = big.NewInt(0)
				}

				out = append(out, &callEntry{
					From:   top.From,
					To:     top.To,
					Value:  v,
					Type:   top.Type,
					Method: top.Method,
					Return: call.Data,
				})
			}
		} else {
			stack.Push(call)
		}
	}

	return out
}
