package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/state"
)

type EthCallReq struct {
	From common.Address `json:"from"`
	To   common.Address `json:"to"`
	Data hexutil.Bytes  `json:"data"`
}

type TraceCallReq struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Data         hexutil.Bytes  `json:"data"`
	MaxFeePerGas hexutil.Big    `json:"maxFeePerGas"`
}

type TraceStateOverrides struct {
	Balance hexutil.Big `json:"balance"`
}

type TraceCallOpts struct {
	Tracer         string            `json:"tracer"`
	StateOverrides state.OverrideSet `json:"stateOverrides"`
}

var (
	// A dummy private key used to build *bind.TransactOpts for simulation.
	DummyPk, _ = crypto.GenerateKey()
)
