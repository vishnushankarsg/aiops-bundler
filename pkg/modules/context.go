package modules

import (
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aimiddleware/stake"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/ethereum/go-ethereum/common"
)

type PendingRemovalItem struct {
	Op     *aiop.AiOperation
	Reason string
}

// BatchHandlerCtx is the object passed to BatchHandler functions during the Bundler's Run process. It
// also contains a Data field for adding arbitrary key-value pairs to the context. These values will be
// logged by the Bundler at the end of each run.
type BatchHandlerCtx struct {
	Batch          []*aiop.AiOperation
	PendingRemoval []*PendingRemovalItem
	AiMiddleware   common.Address
	ChainID        *big.Int
	BaseFee        *big.Int
	Tip            *big.Int
	GasPrice       *big.Int
	Data           map[string]any
}

// NewBatchHandlerContext creates a new BatchHandlerCtx using a copy of the given batch.
func NewBatchHandlerContext(
	batch []*aiop.AiOperation,
	aiMiddleware common.Address,
	chainID *big.Int,
	baseFee *big.Int,
	tip *big.Int,
	gasPrice *big.Int,
) *BatchHandlerCtx {
	var copy []*aiop.AiOperation
	copy = append(copy, batch...)

	return &BatchHandlerCtx{
		Batch:          copy,
		PendingRemoval: []*PendingRemovalItem{},
		AiMiddleware:   aiMiddleware,
		ChainID:        chainID,
		BaseFee:        baseFee,
		Tip:            tip,
		GasPrice:       gasPrice,
		Data:           make(map[string]any),
	}
}

// MarkOpIndexForRemoval will remove the op by index from the batch and add it to the pending removal array.
// This should be used for ops that are not to be included on-chain and dropped from the mempool.
func (c *BatchHandlerCtx) MarkOpIndexForRemoval(index int, reason string) {
	batch := []*aiop.AiOperation{}
	var op *aiop.AiOperation
	for i, curr := range c.Batch {
		if i == index {
			op = curr
		} else {
			batch = append(batch, curr)
		}
	}
	if op == nil {
		return
	}

	c.Batch = batch
	c.PendingRemoval = append(c.PendingRemoval, &PendingRemovalItem{
		Op:     op,
		Reason: reason,
	})
}

// AiOpHandlerCtx is the object passed to AiOpHandler functions during the Client's SendAiOperation
// process.
type AiOpHandlerCtx struct {
	AiOp                *aiop.AiOperation
	AiMiddleware        common.Address
	ChainID             *big.Int
	pendingSenderOps    []*aiop.AiOperation
	pendingFactoryOps   []*aiop.AiOperation
	pendingPaymasterOps []*aiop.AiOperation
	senderDeposit       *aimiddleware.IDepositManagerDepositInfo
	factoryDeposit      *aimiddleware.IDepositManagerDepositInfo
	paymasterDeposit    *aimiddleware.IDepositManagerDepositInfo
}

// NewAiOpHandlerContext creates a new AiOpHandlerCtx using a given op.
func NewAiOpHandlerContext(
	op *aiop.AiOperation,
	aiMiddleware common.Address,
	chainID *big.Int,
	mem *mempool.Mempool,
	gs stake.GetStakeFunc,
) (*AiOpHandlerCtx, error) {
	// Fetch any pending AiOperations in the mempool by entity
	pso, err := mem.GetOps(aiMiddleware, op.Sender)
	if err != nil {
		return nil, err
	}
	pfo, err := mem.GetOps(aiMiddleware, op.GetFactory())
	if err != nil {
		return nil, err
	}
	ppo, err := mem.GetOps(aiMiddleware, op.GetPaymaster())
	if err != nil {
		return nil, err
	}

	// Fetch the current aimiddleware deposits by entity
	sd, err := gs(aiMiddleware, op.Sender)
	if err != nil {
		return nil, err
	}
	fd, err := gs(aiMiddleware, op.GetFactory())
	if err != nil {
		return nil, err
	}
	pd, err := gs(aiMiddleware, op.GetPaymaster())
	if err != nil {
		return nil, err
	}

	return &AiOpHandlerCtx{
		AiOp:                op,
		AiMiddleware:        aiMiddleware,
		ChainID:             chainID,
		pendingSenderOps:    pso,
		pendingFactoryOps:   pfo,
		pendingPaymasterOps: ppo,
		senderDeposit:       sd,
		factoryDeposit:      fd,
		paymasterDeposit:    pd,
	}, nil
}

// GetSenderDepositInfo returns the current AiMiddleware deposit for the sender.
func (c *AiOpHandlerCtx) GetSenderDepositInfo() *aimiddleware.IDepositManagerDepositInfo {
	return c.senderDeposit
}

// GetFactoryDepositInfo returns the current AiMiddleware deposit for the factory.
func (c *AiOpHandlerCtx) GetFactoryDepositInfo() *aimiddleware.IDepositManagerDepositInfo {
	return c.factoryDeposit
}

// GetPaymasterDepositInfo returns the current AiMiddleware deposit for the paymaster.
func (c *AiOpHandlerCtx) GetPaymasterDepositInfo() *aimiddleware.IDepositManagerDepositInfo {
	return c.paymasterDeposit
}

// GetPendingSenderOps returns all pending AiOperations in the mempool by the same sender.
func (c *AiOpHandlerCtx) GetPendingSenderOps() []*aiop.AiOperation {
	return c.pendingSenderOps
}

// GetPendingFactoryOps returns all pending AiOperations in the mempool by the same factory.
func (c *AiOpHandlerCtx) GetPendingFactoryOps() []*aiop.AiOperation {
	return c.pendingFactoryOps
}

// GetPendingPaymasterOps returns all pending AiOperations in the mempool by the same paymaster.
func (c *AiOpHandlerCtx) GetPendingPaymasterOps() []*aiop.AiOperation {
	return c.pendingPaymasterOps
}
