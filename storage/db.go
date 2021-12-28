package storage

import "github.com/dgraph-io/badger/v3"

type Database interface {
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Close() error
}

var (
	ErrKeyNotFound = badger.ErrKeyNotFound
)
