package transaction

import (
	"context"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/reverts"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/signer"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Opts contains all the fields required for submitting a transaction to call HandleOps on the AiMiddleware
// contract.
type Opts struct {
	// Options for the network
	EOA     *signer.EOA
	Eth     *ethclient.Client
	ChainID *big.Int

	// Options for the AiMiddleware
	AiMiddleware common.Address
	Batch        []*aiop.AiOperation
	Beneficiary  common.Address

	// Options for the EOA transaction
	BaseFee     *big.Int
	Tip         *big.Int
	GasPrice    *big.Int
	GasLimit    uint64
	NoSend      bool
	WaitTimeout time.Duration
}

func toAbiType(batch []*aiop.AiOperation) []aimiddleware.AiOperation {
	ops := []aimiddleware.AiOperation{}
	for _, op := range batch {
		ops = append(ops, aimiddleware.AiOperation(*op))
	}

	return ops
}

// EstimateHandleOpsGas returns a gas estimate required to call handleOps() with a given batch. A failed call
// will return the cause of the revert.
func EstimateHandleOpsGas(opts *Opts) (gas uint64, revert *reverts.FailedOpRevert, err error) {
	ep, err := aimiddleware.NewAimiddleware(opts.AiMiddleware, opts.Eth)
	if err != nil {
		return 0, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(opts.EOA.PrivateKey, opts.ChainID)
	if err != nil {
		return 0, nil, err
	}
	auth.GasLimit = math.MaxUint64
	auth.NoSend = true

	tx, err := ep.HandleOps(auth, toAbiType(opts.Batch), opts.Beneficiary)
	if err != nil {
		return 0, nil, err
	}

	est, err := opts.Eth.EstimateGas(context.Background(), ethereum.CallMsg{
		From:       opts.EOA.Address,
		To:         tx.To(),
		Gas:        tx.Gas(),
		GasPrice:   tx.GasPrice(),
		GasFeeCap:  tx.GasFeeCap(),
		GasTipCap:  tx.GasTipCap(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	})
	if err != nil {
		revert, err := reverts.NewFailedOp(err)
		if err != nil {
			return 0, nil, err
		}
		return 0, revert, nil
	}

	return est, nil, nil
}

// HandleOps submits a transaction to send a batch of AiOperations to the AiMiddleware.
func HandleOps(opts *Opts) (txn *types.Transaction, err error) {
	ep, err := aimiddleware.NewAimiddleware(opts.AiMiddleware, opts.Eth)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(opts.EOA.PrivateKey, opts.ChainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = opts.GasLimit
	auth.NoSend = opts.NoSend

	nonce, err := opts.Eth.NonceAt(context.Background(), opts.EOA.Address, nil)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	if opts.BaseFee != nil && opts.Tip != nil {
		auth.GasTipCap = SuggestMeanGasTipCap(opts.Tip, opts.Batch)
		auth.GasFeeCap = SuggestMeanGasFeeCap(opts.BaseFee, opts.Tip, opts.Batch)
	} else if opts.GasPrice != nil {
		auth.GasPrice = SuggestMeanGasPrice(opts.GasPrice, opts.Batch)
	} else {
		return nil, errors.New("transaction: either the dynamic or legacy gas fees must be set")
	}

	txn, err = ep.HandleOps(auth, toAbiType(opts.Batch), opts.Beneficiary)
	if err != nil {
		return nil, err
	} else if opts.WaitTimeout == 0 || opts.NoSend {
		// Don't wait for transaction to be included. All aiOps in the current batch will be dropped
		// regardless of the transaction status.
		return txn, nil
	}

	return Wait(txn, opts.Eth, opts.WaitTimeout)
}
