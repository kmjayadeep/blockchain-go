package storage

import (
	"github.com/dgraph-io/badger/v3"
)

type DB struct {
	db *badger.DB
}

func NewDatabase(opts badger.Options) (*DB, error) {
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &DB{
		db,
	}, nil
}

func (d *DB) Get(key string) ([]byte, error) {
	var result []byte
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		result = valCopy
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DB) Put(key string, val []byte) error {
	err := d.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		return err
	})

	return err
}

func (d *DB) Close() error {
	return d.db.Close()
}
