// Package mempool provides a local representation of all the AiOperations that are known to the bundler
// which have passed all Client checks and pending action by the Bundler.
package mempool

import (
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
)

// Mempool provides read and write access to a pool of pending AiOperations which have passed all Client
// checks.
type Mempool struct {
	db    *badger.DB
	queue *aiOpQueues
}

// New creates an instance of a mempool that uses an embedded DB to persist and load AiOperations from disk
// incase of a reset.
func New(db *badger.DB) (*Mempool, error) {
	queue := newAiOpQueue()
	err := loadFromDisk(db, queue)
	if err != nil {
		return nil, err
	}

	return &Mempool{db, queue}, nil
}

// GetOps returns all the AiOperations associated with an AiMiddleware and Sender address.
func (m *Mempool) GetOps(aiMiddleware common.Address, sender common.Address) ([]*aiop.AiOperation, error) {
	ops := m.queue.GetOps(aiMiddleware, sender)
	return ops, nil
}

// AddOp adds a AiOperation to the mempool or replace an existing one with the same AiMiddleware, Sender, and
// Nonce values.
func (m *Mempool) AddOp(aiMiddleware common.Address, op *aiop.AiOperation) error {
	data, err := op.MarshalJSON()
	if err != nil {
		return err
	}

	err = m.db.Update(func(txn *badger.Txn) error {
		return txn.Set(getUniqueKey(aiMiddleware, op.Sender, op.Nonce), data)
	})
	if err != nil {
		return err
	}

	m.queue.AddOp(aiMiddleware, op)
	return nil
}

// RemoveOps removes a list of AiOperations from the mempool by AiMiddleware, Sender, and Nonce values.
func (m *Mempool) RemoveOps(aiMiddleware common.Address, ops ...*aiop.AiOperation) error {
	err := m.db.Update(func(txn *badger.Txn) error {
		for _, op := range ops {
			err := txn.Delete(getUniqueKey(aiMiddleware, op.Sender, op.Nonce))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	m.queue.RemoveOps(aiMiddleware, ops...)
	return nil
}

// Dump will return a list of AiOperations from the mempool by AiMiddleware in the order it arrived.
func (m *Mempool) Dump(aiMiddleware common.Address) ([]*aiop.AiOperation, error) {
	return m.queue.All(aiMiddleware), nil
}

// Clear will clear the entire embedded db and reset it to a clean state.
func (m *Mempool) Clear() error {
	if err := m.db.DropAll(); err != nil {
		return err
	}
	m.queue = newAiOpQueue()

	return nil
}
