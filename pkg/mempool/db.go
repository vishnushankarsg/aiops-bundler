package mempool

import (
	"encoding/json"
	"math/big"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/dbutils"
	"github.com/DAO-Metaplayer/aiops-bundler/pkg/aiop"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
)

var (
	keyPrefix = dbutils.JoinValues("mempool")
)

func getUniqueKey(aiMiddleware common.Address, sender common.Address, nonce *big.Int) []byte {
	return []byte(
		dbutils.JoinValues(keyPrefix, aiMiddleware.String(), sender.String(), nonce.String()),
	)
}

func getAiMiddlewareFromDBKey(key []byte) common.Address {
	slc := dbutils.SplitValues(string(key))
	return common.HexToAddress(slc[1])
}

func getAiOpFromDBValue(value []byte) (*aiop.AiOperation, error) {
	data := make(map[string]any)
	if err := json.Unmarshal(value, &data); err != nil {
		return nil, err
	}

	op, err := aiop.New(data)
	if err != nil {
		return nil, err
	}

	return op, nil
}

func loadFromDisk(db *badger.DB, q *aiOpQueues) error {
	return db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		prefix := []byte(keyPrefix)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			ep := getAiMiddlewareFromDBKey(item.Key())

			err := item.Value(func(v []byte) error {
				op, err := getAiOpFromDBValue(v)
				if err != nil {
					return err
				}

				q.AddOp(ep, op)
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})
}
