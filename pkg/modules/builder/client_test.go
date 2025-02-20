package builder

import (
	"errors"
	"math/big"
	"testing"

	"github.com/AO-Metaplayer/aiops-bundler/internal/testutils"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/metachris/flashbotsrpc"
)

func TestSendAiOperationWithAllUpstreamErrors(t *testing.T) {
	n := testutils.RpcMock(testutils.MethodMocks{
		"eth_blockNumber":           "0x1",
		"eth_gasPrice":              "0x1",
		"eth_getTransactionCount":   "0x1",
		"eth_estimateGas":           "0x1",
		"eth_getBlockByNumber":      testutils.NewBlockMock(),
		"eth_getTransactionReceipt": testutils.NewTransactionReceiptMock(),
	})
	r, _ := rpc.Dial(n.URL)
	eth := ethclient.NewClient(r)

	bb1 := testutils.BadBuilderRpcMock()
	bb2 := testutils.BadBuilderRpcMock()
	fb := flashbotsrpc.NewBuilderBroadcastRPC([]string{bb1.URL, bb2.URL})
	fn := New(testutils.DummyEOA, eth, fb, testutils.DummyEOA.Address, 1).SendAiOperation()

	if err := fn(
		modules.NewBatchHandlerContext(
			[]*aiop.AiOperation{testutils.MockValidInitAiOp()},
			common.HexToAddress("0x"),
			testutils.ChainID,
			big.NewInt(1),
			big.NewInt(1),
			big.NewInt(1),
		),
	); !errors.Is(err, ErrFlashbotsBroadcastBundle) {
		t.Fatalf("got %v, want ErrFlashbotsBroadcastBundle", err)
	}
}

func TestSendAiOperationWithPartialUpstreamErrors(t *testing.T) {
	n := testutils.RpcMock(testutils.MethodMocks{
		"eth_blockNumber":           "0x1",
		"eth_gasPrice":              "0x1",
		"eth_getTransactionCount":   "0x1",
		"eth_estimateGas":           "0x1",
		"eth_getBlockByNumber":      testutils.NewBlockMock(),
		"eth_getTransactionReceipt": testutils.NewTransactionReceiptMock(),
	})
	r, _ := rpc.Dial(n.URL)
	eth := ethclient.NewClient(r)

	bb1 := testutils.RpcMock(testutils.MethodMocks{
		"eth_sendBundle": map[string]string{
			"bundleHash": testutils.MockHash,
		},
	})
	bb2 := testutils.BadBuilderRpcMock()
	fb := flashbotsrpc.NewBuilderBroadcastRPC([]string{bb1.URL, bb2.URL})
	fn := New(testutils.DummyEOA, eth, fb, testutils.DummyEOA.Address, 1).SendAiOperation()

	if err := fn(
		modules.NewBatchHandlerContext(
			[]*aiop.AiOperation{testutils.MockValidInitAiOp()},
			common.HexToAddress("0x"),
			testutils.ChainID,
			big.NewInt(1),
			big.NewInt(1),
			big.NewInt(1),
		),
	); err != nil {
		t.Fatalf("got %v, want nil", err)
	}
}

func TestSendAiOperationWithNoUpstreamErrors(t *testing.T) {
	n := testutils.RpcMock(testutils.MethodMocks{
		"eth_blockNumber":           "0x1",
		"eth_gasPrice":              "0x1",
		"eth_getTransactionCount":   "0x1",
		"eth_estimateGas":           "0x1",
		"eth_getBlockByNumber":      testutils.NewBlockMock(),
		"eth_getTransactionReceipt": testutils.NewTransactionReceiptMock(),
	})
	r, _ := rpc.Dial(n.URL)
	eth := ethclient.NewClient(r)

	bb1 := testutils.RpcMock(testutils.MethodMocks{
		"eth_sendBundle": map[string]string{
			"bundleHash": testutils.MockHash,
		},
	})
	bb2 := testutils.RpcMock(testutils.MethodMocks{
		"eth_sendBundle": map[string]string{
			"bundleHash": testutils.MockHash,
		},
	})
	fb := flashbotsrpc.NewBuilderBroadcastRPC([]string{bb1.URL, bb2.URL})
	fn := New(testutils.DummyEOA, eth, fb, testutils.DummyEOA.Address, 1).SendAiOperation()

	if err := fn(
		modules.NewBatchHandlerContext(
			[]*aiop.AiOperation{testutils.MockValidInitAiOp()},
			common.HexToAddress("0x"),
			testutils.ChainID,
			big.NewInt(1),
			big.NewInt(1),
			big.NewInt(1),
		),
	); err != nil {
		t.Fatalf("got %v, want nil", err)
	}
}
