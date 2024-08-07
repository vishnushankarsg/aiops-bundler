package client

import (
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/filter"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/fees"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// GetAiOpReceiptFunc is a general interface for fetching a AiOperationReceipt given a aiOpHash,
// AiMiddleware address, and block range.
type GetAiOpReceiptFunc = func(hash string, ep common.Address, blkRange uint64) (*filter.AiOperationReceipt, error)

func getAiOpReceiptNoop() GetAiOpReceiptFunc {
	return func(hash string, ep common.Address, blkRange uint64) (*filter.AiOperationReceipt, error) {
		return nil, nil
	}
}

// GetAiOpReceiptWithEthClient returns an implementation of GetAiOpReceiptFunc that relies on an eth
// client to fetch a AiOperationReceipt.
func GetAiOpReceiptWithEthClient(eth *ethclient.Client) GetAiOpReceiptFunc {
	return func(hash string, ep common.Address, blkRange uint64) (*filter.AiOperationReceipt, error) {
		return filter.GetAiOperationReceipt(eth, hash, ep, blkRange)
	}
}

// GetGasPricesFunc is a general interface for fetching values for maxFeePerGas and maxPriorityFeePerGas.
type GetGasPricesFunc = func() (*fees.GasPrices, error)

func getGasPricesNoop() GetGasPricesFunc {
	return func() (*fees.GasPrices, error) {
		return &fees.GasPrices{
			MaxFeePerGas:         big.NewInt(0),
			MaxPriorityFeePerGas: big.NewInt(0),
		}, nil
	}
}

// GetGasPricesWithEthClient returns an implementation of GetGasPricesFunc that relies on an eth client to
// fetch values for maxFeePerGas and maxPriorityFeePerGas.
func GetGasPricesWithEthClient(eth *ethclient.Client) GetGasPricesFunc {
	return func() (*fees.GasPrices, error) {
		return fees.NewGasPrices(eth)
	}
}

// GetGasEstimateFunc is a general interface for fetching an estimate for verificationGasLimit and
// callGasLimit given a aiOp and AiMiddleware address.
type GetGasEstimateFunc = func(
	ep common.Address,
	op *aiop.AiOperation,
	sos state.OverrideSet,
) (verificationGas uint64, callGas uint64, err error)

func getGasEstimateNoop() GetGasEstimateFunc {
	return func(
		ep common.Address,
		op *aiop.AiOperation,
		sos state.OverrideSet,
	) (verificationGas uint64, callGas uint64, err error) {
		return 0, 0, nil
	}
}

// GetGasEstimateWithEthClient returns an implementation of GetGasEstimateFunc that relies on an eth client to
// fetch an estimate for verificationGasLimit and callGasLimit.
func GetGasEstimateWithEthClient(
	rpc *rpc.Client,
	ov *gas.Overhead,
	chain *big.Int,
	maxGasLimit *big.Int,
	tracer string,
) GetGasEstimateFunc {
	return func(
		ep common.Address,
		op *aiop.AiOperation,
		sos state.OverrideSet,
	) (verificationGas uint64, callGas uint64, err error) {
		return gas.EstimateGas(&gas.EstimateInput{
			Rpc:          rpc,
			AiMiddleware: ep,
			Op:           op,
			Sos:          sos,
			Ov:           ov,
			ChainID:      chain,
			MaxGasLimit:  maxGasLimit,
			Tracer:       tracer,
		})
	}
}

// GetAiOpByHashFunc is a general interface for fetching a AiOperation given a aiOpHash, AiMiddleware
// address, chain ID, and block range.
type GetAiOpByHashFunc func(hash string, ep common.Address, chain *big.Int, blkRange uint64) (*filter.HashLookupResult, error)

func getAiOpByHashNoop() GetAiOpByHashFunc {
	return func(hash string, ep common.Address, chain *big.Int, blkRange uint64) (*filter.HashLookupResult, error) {
		return nil, nil
	}
}

// GetAiOpByHashWithEthClient returns an implementation of GetAiOpByHashFunc that relies on an eth client
// to fetch a AiOperation.
func GetAiOpByHashWithEthClient(eth *ethclient.Client) GetAiOpByHashFunc {
	return func(hash string, ep common.Address, chain *big.Int, blkRange uint64) (*filter.HashLookupResult, error) {
		return filter.GetAiOperationByHash(eth, hash, ep, chain, blkRange)
	}
}
