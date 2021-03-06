package storage_test

import (
	"bytes"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/storage"
)

func TestDB(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()

	testKey := "testkey"
	testVal := []byte("testval")

	err = db.Put(testKey, testVal)

	if err != nil {
		t.Errorf("unable to put data")
	}

	val, err := db.Get(testKey)

	if err != nil {
		t.Errorf("unable to get data, error : %s", err.Error())
	}

	if bytes.Compare(val, testVal) != 0 {
		t.Errorf("unable to get data")
	}

}
