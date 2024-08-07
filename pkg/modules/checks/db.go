package checks

import (
	"encoding/json"

	"github.com/DAO-Metaplayer/aiops-bundler/internal/dbutils"
	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
)

var (
	keyPrefix        = dbutils.JoinValues("checks")
	codeHashesPrefix = dbutils.JoinValues(keyPrefix, "codeHashes")
)

func getCodeHashesKey(aiOpHash common.Hash) []byte {
	return []byte(dbutils.JoinValues(codeHashesPrefix, aiOpHash.String()))
}

func saveCodeHashes(db *badger.DB, aiOpHash common.Hash, codeHashes []codeHash) error {
	return db.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(codeHashes)
		if err != nil {
			return err
		}

		return txn.Set(getCodeHashesKey(aiOpHash), data)
	})
}

func getSavedCodeHashes(db *badger.DB, aiOpHash common.Hash) ([]codeHash, error) {
	var ch []codeHash
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(getCodeHashesKey(aiOpHash))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &ch)
		})
	})

	return ch, err
}

func removeSavedCodeHashes(db *badger.DB, aiOpHashes ...common.Hash) error {
	return db.Update(func(txn *badger.Txn) error {
		for _, aiOpHash := range aiOpHashes {
			if err := txn.Delete(getCodeHashesKey(aiOpHash)); err != nil {
				return err
			}
		}
		return nil
	})
}
