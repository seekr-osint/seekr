package api

import (
	"encoding/json"
	"fmt"
	badger "github.com/dgraph-io/badger/v4"
)

func DefaultSaveDB(config ApiConfig) error {
	db, err := badger.Open(badger.DefaultOptions(config.DataBaseFile))
	if err != nil {
		return fmt.Errorf("failed to open badger db: %w. DataBaseFile: %s", err, config.DataBaseFile)
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		jsonBytes, err := json.MarshalIndent(config.DataBase, "", "\t")
		if err != nil {
			return fmt.Errorf("error marshaling database to JSON: %w", err)
		}
		return txn.Set([]byte("data"), jsonBytes)
	})
	if err != nil {
		return fmt.Errorf("error saving database to BadgerDB: %w", err)
	}

	return nil
}

func DefaultLoadDB(config ApiConfig) (ApiConfig, error) {

	db, err := badger.Open(badger.DefaultOptions(config.DataBaseFile))
	if err != nil {
		return config, fmt.Errorf("failed to open badger db: %w. DataBaseFile: %s", err, config.DataBaseFile)
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("data"))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				// if key is not found, return empty config
				return nil
			}
			return fmt.Errorf("error getting item from BadgerDB: %w", err)
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &config.DataBase)
		})
		if err != nil {
			return fmt.Errorf("error unmarshaling data from BadgerDB: %w", err)
		}

		return nil
	})
	if err != nil {
		return config, fmt.Errorf("error loading database from BadgerDB: %w", err)
	}

	return config, nil
}
