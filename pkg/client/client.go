// Package client provides the mediator for processing incoming AiOperations to the bundler.
package client

import (
	"errors"
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/logger"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/filter"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/stake"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/modules/noop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-logr/logr"
)

// Client controls the end to end process of adding incoming AiOperations to the mempool. It also
// implements the required RPC methods as specified in EIP-4337.
type Client struct {
	mempool                *mempool.Mempool
	ov                     *gas.Overhead
	chainID                *big.Int
	supportedAiMiddlewares []common.Address
	aiOpHandler            modules.AiOpHandlerFunc
	logger                 logr.Logger
	getAiOpReceipt         GetAiOpReceiptFunc
	getGasPrices           GetGasPricesFunc
	getGasEstimate         GetGasEstimateFunc
	getAiOpByHash          GetAiOpByHashFunc
	getStakeFunc           stake.GetStakeFunc
	opLookupLimit          uint64
}

// New initializes a new ERC-4337 client which can be extended with modules for validating AiOperations
// that are allowed to be added to the mempool.
func New(
	mempool *mempool.Mempool,
	ov *gas.Overhead,
	chainID *big.Int,
	supportedAiMiddlewares []common.Address,
	opLookupLimit uint64,
) *Client {
	return &Client{
		mempool:                mempool,
		ov:                     ov,
		chainID:                chainID,
		supportedAiMiddlewares: supportedAiMiddlewares,
		aiOpHandler:            noop.AiOpHandler,
		logger:                 logger.NewZeroLogr().WithName("client"),
		getAiOpReceipt:         getAiOpReceiptNoop(),
		getGasPrices:           getGasPricesNoop(),
		getGasEstimate:         getGasEstimateNoop(),
		getAiOpByHash:          getAiOpByHashNoop(),
		getStakeFunc:           stake.GetStakeFuncNoop(),
		opLookupLimit:          opLookupLimit,
	}
}

func (i *Client) parseAiMiddlewareAddress(ep string) (common.Address, error) {
	for _, addr := range i.supportedAiMiddlewares {
		if common.HexToAddress(ep) == addr {
			return addr, nil
		}
	}

	return common.Address{}, errors.New("aiMiddleware: Implementation not supported")
}

// UseLogger defines the logger object used by the Client instance based on the go-logr/logr interface.
func (i *Client) UseLogger(logger logr.Logger) {
	i.logger = logger.WithName("client")
}

// UseModules defines the AiOpHandlers to process a aiOp after it has gone through the standard checks.
func (i *Client) UseModules(handlers ...modules.AiOpHandlerFunc) {
	i.aiOpHandler = modules.ComposeAiOpHandlerFunc(handlers...)
}

// SetGetAiOpReceiptFunc defines a general function for fetching a AiOpReceipt given a aiOpHash and
// AiMiddleware address. This function is called in *Client.GetAiOperationReceipt.
func (i *Client) SetGetAiOpReceiptFunc(fn GetAiOpReceiptFunc) {
	i.getAiOpReceipt = fn
}

// SetGetGasPricesFunc defines a general function for fetching values for maxFeePerGas and
// maxPriorityFeePerGas. This function is called in *Client.EstimateAiOperationGas if given fee values are
// 0.
func (i *Client) SetGetGasPricesFunc(fn GetGasPricesFunc) {
	i.getGasPrices = fn
}

// SetGetGasEstimateFunc defines a general function for fetching an estimate for verificationGasLimit and
// callGasLimit given a aiOp and AiMiddleware address. This function is called in
// *Client.EstimateAiOperationGas.
func (i *Client) SetGetGasEstimateFunc(fn GetGasEstimateFunc) {
	i.getGasEstimate = fn
}

// SetGetAiOpByHashFunc defines a general function for fetching a aiOp given a aiOpHash, AiMiddleware
// address, and chain ID. This function is called in *Client.GetAiOperationByHash.
func (i *Client) SetGetAiOpByHashFunc(fn GetAiOpByHashFunc) {
	i.getAiOpByHash = fn
}

// SetGetStakeFunc defines a general function for retrieving the AiMiddleware stake for a given address. This
// function is called in *Client.SendAiOperation to create a context.
func (i *Client) SetGetStakeFunc(fn stake.GetStakeFunc) {
	i.getStakeFunc = fn
}

// SendAiOperation implements the method call for eth_sendAiOperation.
// It returns true if aiOp was accepted otherwise returns an error.
func (i *Client) SendAiOperation(op map[string]any, ep string) (string, error) {
	// Init logger
	l := i.logger.WithName("eth_sendAiOperation")

	// Check AiMiddleware and aiOp is valid.
	epAddr, err := i.parseAiMiddlewareAddress(ep)
	if err != nil {
		l.Error(err, "eth_sendAiOperation error")
		return "", err
	}
	l = l.
		WithValues("aimiddleware", epAddr.String()).
		WithValues("chain_id", i.chainID.String())

	aiOp, err := aiop.New(op)
	if err != nil {
		l.Error(err, "eth_sendAiOperation error")
		return "", err
	}
	hash := aiOp.GetAiOpHash(epAddr, i.chainID)
	l = l.WithValues("aiop_hash", hash)

	// Run through client module stack.
	ctx, err := modules.NewAiOpHandlerContext(
		aiOp,
		epAddr,
		i.chainID,
		i.mempool,
		i.getStakeFunc,
	)
	if err != nil {
		l.Error(err, "eth_sendAiOperation error")
		return "", err
	}
	if err := i.aiOpHandler(ctx); err != nil {
		l.Error(err, "eth_sendAiOperation error")
		return "", err
	}

	// Add aiOp to mempool.
	if err := i.mempool.AddOp(epAddr, ctx.AiOp); err != nil {
		l.Error(err, "eth_sendAiOperation error")
		return "", err
	}

	l.Info("eth_sendAiOperation ok")
	return hash.String(), nil
}

// EstimateAiOperationGas returns estimates for PreVerificationGas, VerificationGasLimit, and CallGasLimit
// given a AiOperation, AiMiddleware address, and state OverrideSet. The signature field and current gas
// values will not be validated although there should be dummy values in place for the most reliable results
// (e.g. a signature with the correct length).
func (i *Client) EstimateAiOperationGas(
	op map[string]any,
	ep string,
	os map[string]any,
) (*gas.GasEstimates, error) {
	// Init logger
	l := i.logger.WithName("eth_estimateAiOperationGas")

	// Check AiMiddleware and aiOp is valid.
	epAddr, err := i.parseAiMiddlewareAddress(ep)
	if err != nil {
		l.Error(err, "eth_estimateAiOperationGas error")
		return nil, err
	}
	l = l.
		WithValues("aimiddleware", epAddr.String()).
		WithValues("chain_id", i.chainID.String())

	aiOp, err := aiop.New(op)
	if err != nil {
		l.Error(err, "eth_estimateAiOperationGas error")
		return nil, err
	}
	hash := aiOp.GetAiOpHash(epAddr, i.chainID)
	l = l.WithValues("aiop_hash", hash)

	// Parse state override set.
	sos, err := state.ParseOverrideData(os)
	if err != nil {
		l.Error(err, "eth_estimateAiOperationGas error")
		return nil, err
	}

	// Override op with suggested gas prices if maxFeePerGas is 0. This allows for more reliable gas
	// estimations upstream. The default balance override also ensures simulations won't revert on
	// insufficient funds.
	if aiOp.MaxFeePerGas.Cmp(common.Big0) != 1 {
		gp, err := i.getGasPrices()
		if err != nil {
			l.Error(err, "eth_estimateAiOperationGas error")
			return nil, err
		}
		aiOp.MaxFeePerGas = gp.MaxFeePerGas
		aiOp.MaxPriorityFeePerGas = gp.MaxPriorityFeePerGas
	}

	// Estimate gas limits
	vg, cg, err := i.getGasEstimate(epAddr, aiOp, sos)
	if err != nil {
		l.Error(err, "eth_estimateAiOperationGas error")
		return nil, err
	}

	// Calculate PreVerificationGas
	pvg, err := i.ov.CalcPreVerificationGasWithBuffer(aiOp)
	if err != nil {
		l.Error(err, "eth_estimateAiOperationGas error")
		return nil, err
	}

	l.Info("eth_estimateAiOperationGas ok")
	return &gas.GasEstimates{
		PreVerificationGas:   pvg,
		VerificationGasLimit: big.NewInt(int64(vg)),
		CallGasLimit:         big.NewInt(int64(cg)),

		// TODO: Deprecate in v0.7
		VerificationGas: big.NewInt(int64(vg)),
	}, nil
}

// GetAiOperationReceipt fetches a AiOperation receipt based on a aiOpHash returned by
// *Client.SendAiOperation.
func (i *Client) GetAiOperationReceipt(
	hash string,
) (*filter.AiOperationReceipt, error) {
	// Init logger
	l := i.logger.WithName("eth_getAiOperationReceipt").WithValues("aiop_hash", hash)

	ev, err := i.getAiOpReceipt(hash, i.supportedAiMiddlewares[0], i.opLookupLimit)
	if err != nil {
		l.Error(err, "eth_getAiOperationReceipt error")
		return nil, err
	}

	l.Info("eth_getAiOperationReceipt ok")
	return ev, nil
}

// GetAiOperationByHash returns a AiOperation based on a given aiOpHash returned by
// *Client.SendAiOperation.
func (i *Client) GetAiOperationByHash(hash string) (*filter.HashLookupResult, error) {
	// Init logger
	l := i.logger.WithName("eth_getAiOperationByHash").WithValues("aiop_hash", hash)

	res, err := i.getAiOpByHash(hash, i.supportedAiMiddlewares[0], i.chainID, i.opLookupLimit)
	if err != nil {
		l.Error(err, "eth_getAiOperationByHash error")
		return nil, err
	}

	return res, nil
}

// SupportedAiMiddlewares implements the method call for eth_supportedAiMiddlewares. It returns the array of
// AiMiddleware addresses that is supported by the client. The first address in the array is the preferred
// AiMiddleware.
func (i *Client) SupportedAiMiddlewares() ([]string, error) {
	slc := []string{}
	for _, ep := range i.supportedAiMiddlewares {
		slc = append(slc, ep.String())
	}

	return slc, nil
}

// ChainID implements the method call for eth_chainId. It returns the current chainID used by the client.
// This method is used to validate that the client's chainID is in sync with the caller.
func (i *Client) ChainID() (string, error) {
	return hexutil.EncodeBig(i.chainID), nil
}
