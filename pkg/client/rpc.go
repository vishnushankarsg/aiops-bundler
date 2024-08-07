package client

import (
	"errors"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/filter"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/gas"
)

// Named AiOperation type for jsonrpc package.
type aiOperation map[string]any

// Named StateOverride type for jsonrpc package.
type optional_stateOverride map[string]any

// RpcAdapter is an adapter for routing JSON-RPC method calls to the correct client functions.
type RpcAdapter struct {
	client *Client
	debug  *Debug
}

// NewRpcAdapter initializes a new RpcAdapter which can be used with a JSON-RPC server.
func NewRpcAdapter(client *Client, debug *Debug) *RpcAdapter {
	return &RpcAdapter{client, debug}
}

// Eth_sendAiOperation routes method calls to *Client.SendAiOperation.
func (r *RpcAdapter) Eth_sendAiOperation(op aiOperation, ep string) (string, error) {
	return r.client.SendAiOperation(op, ep)
}

// Eth_estimateAiOperationGas routes method calls to *Client.EstimateAiOperationGas.
func (r *RpcAdapter) Eth_estimateAiOperationGas(
	op aiOperation,
	ep string,
	os optional_stateOverride,
) (*gas.GasEstimates, error) {
	return r.client.EstimateAiOperationGas(op, ep, os)
}

// Eth_getAiOperationReceipt routes method calls to *Client.GetAiOperationReceipt.
func (r *RpcAdapter) Eth_getAiOperationReceipt(
	aiOpHash string,
) (*filter.AiOperationReceipt, error) {
	return r.client.GetAiOperationReceipt(aiOpHash)
}

// Eth_getAiOperationByHash routes method calls to *Client.GetAiOperationByHash.
func (r *RpcAdapter) Eth_getAiOperationByHash(
	aiOpHash string,
) (*filter.HashLookupResult, error) {
	return r.client.GetAiOperationByHash(aiOpHash)
}

// Eth_supportedAiMiddlewares routes method calls to *Client.SupportedAiMiddlewares.
func (r *RpcAdapter) Eth_supportedAiMiddlewares() ([]string, error) {
	return r.client.SupportedAiMiddlewares()
}

// Eth_chainId routes method calls to *Client.ChainID.
func (r *RpcAdapter) Eth_chainId() (string, error) {
	return r.client.ChainID()
}

// Debug_bundler_clearState routes method calls to *Debug.ClearState.
func (r *RpcAdapter) Debug_bundler_clearState() (string, error) {
	if r.debug == nil {
		return "", errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.ClearState()
}

// Debug_bundler_dumpMempool routes method calls to *Debug.DumpMempool.
func (r *RpcAdapter) Debug_bundler_dumpMempool(ep string) ([]map[string]any, error) {
	if r.debug == nil {
		return []map[string]any{}, errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.DumpMempool(ep)
}

// Debug_bundler_sendBundleNow routes method calls to *Debug.SendBundleNow.
func (r *RpcAdapter) Debug_bundler_sendBundleNow() (string, error) {
	if r.debug == nil {
		return "", errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.SendBundleNow()
}

// Debug_bundler_setBundlingMode routes method calls to *Debug.SetBundlingMode.
func (r *RpcAdapter) Debug_bundler_setBundlingMode(mode string) (string, error) {
	if r.debug == nil {
		return "", errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.SetBundlingMode(mode)
}

// Debug_bundler_setReputation routes method calls to *Debug.SetReputation.
func (r *RpcAdapter) Debug_bundler_setReputation(entries []any, ep string) (string, error) {
	if r.debug == nil {
		return "", errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.SetReputation(entries, ep)
}

// Debug_bundler_dumpReputation routes method calls to *Debug.DumpReputation.
func (r *RpcAdapter) Debug_bundler_dumpReputation(ep string) ([]map[string]any, error) {
	if r.debug == nil {
		return []map[string]any{}, errors.New("rpc: debug mode is not enabled")
	}

	return r.debug.DumpReputation(ep)
}
