package filter

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type parsedTransaction struct {
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       string         `json:"blockNumber"`
	From              common.Address `json:"from"`
	CumulativeGasUsed string         `json:"cumulativeGasUsed"`
	GasUsed           string         `json:"gasUsed"`
	Logs              []*types.Log   `json:"logs"`
	LogsBloom         types.Bloom    `json:"logsBloom"`
	TransactionHash   common.Hash    `json:"transactionHash"`
	TransactionIndex  string         `json:"transactionIndex"`
	EffectiveGasPrice string         `json:"effectiveGasPrice"`
}

type AiOperationReceipt struct {
	AiOpHash      common.Hash        `json:"aiOpHash"`
	Sender        common.Address     `json:"sender"`
	Paymaster     common.Address     `json:"paymaster"`
	Nonce         string             `json:"nonce"`
	Success       bool               `json:"success"`
	ActualGasCost string             `json:"actualGasCost"`
	ActualGasUsed string             `json:"actualGasUsed"`
	From          common.Address     `json:"from"`
	Receipt       *parsedTransaction `json:"receipt"`
	Logs          []*types.Log       `json:"logs"`
}

// GetAiOperationReceipt filters the AiMiddleware contract for AiOperationEvents and returns a receipt for
// both the AiOperation and accompanying transaction.
func GetAiOperationReceipt(
	eth *ethclient.Client,
	aiOpHash string,
	aiMiddleware common.Address,
	blkRange uint64,
) (*AiOperationReceipt, error) {
	if !IsValidAiOpHash(aiOpHash) {
		//lint:ignore ST1005 This needs to match the bundler test spec.
		return nil, errors.New("Missing/invalid aiOpHash")
	}

	it, err := filterAiOperationEvent(eth, aiOpHash, aiMiddleware, blkRange)
	if err != nil {
		return nil, err
	}

	if it.Next() {
		receipt, err := eth.TransactionReceipt(context.Background(), it.Event.Raw.TxHash)
		if err != nil {
			return nil, err
		}
		tx, isPending, err := eth.TransactionByHash(context.Background(), it.Event.Raw.TxHash)
		if err != nil {
			return nil, err
		} else if isPending {
			return nil, nil
		}
		from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			return nil, err
		}

		txnReceipt := &parsedTransaction{
			BlockHash:         receipt.BlockHash,
			BlockNumber:       hexutil.EncodeBig(receipt.BlockNumber),
			From:              from,
			CumulativeGasUsed: hexutil.EncodeBig(big.NewInt(0).SetUint64(receipt.CumulativeGasUsed)),
			GasUsed:           hexutil.EncodeBig(big.NewInt(0).SetUint64(receipt.GasUsed)),
			Logs:              receipt.Logs,
			LogsBloom:         receipt.Bloom,
			TransactionHash:   receipt.TxHash,
			TransactionIndex:  hexutil.EncodeBig(big.NewInt(0).SetUint64(uint64(receipt.TransactionIndex))),
			EffectiveGasPrice: hexutil.EncodeBig(tx.GasPrice()),
		}
		return &AiOperationReceipt{
			AiOpHash:      it.Event.AiOpHash,
			Sender:        it.Event.Sender,
			Paymaster:     it.Event.Paymaster,
			Nonce:         hexutil.EncodeBig(it.Event.Nonce),
			Success:       it.Event.Success,
			ActualGasCost: hexutil.EncodeBig(it.Event.ActualGasCost),
			ActualGasUsed: hexutil.EncodeBig(it.Event.ActualGasUsed),
			From:          from,
			Receipt:       txnReceipt,
			Logs:          []*types.Log{&it.Event.Raw},
		}, nil
	}

	return nil, nil
}
