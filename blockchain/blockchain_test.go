package blockchain

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/storage"
)

func TestInitBlockChain(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()
	chain, err := InitBlockChain(db)

	if chain == nil {
		t.Errorf("unable to initialize chain")
	}

	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	genesisHash := "00031a02a972efd4fa6ea999407149b85b03ccecb8c2bb8eb5a1d068862309d0"
	if fmt.Sprintf("%x", chain.LastHash) != genesisHash {
		t.Errorf("blockchain doesn't have genesis, got hash %x", chain.LastHash)
	}
}

func TestAddBlock(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()
	chain, err := InitBlockChain(db)
	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	chain.AddBlock("testing")
}
