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

	var blocks []*Block
	iter := chain.Iterator()
	block := iter.Next()

	for block != nil {
		blocks = append(blocks, block)
		block = iter.Next()
	}

	if len(blocks) != 1 {
		t.Errorf("blockchain doesn't have genesis, got size %d", len(blocks))
	}

	if string(blocks[0].Data) != "Genesis" {
		t.Errorf("blockchain doesn't have genesis")
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

	err = chain.AddBlock("testing")
	if err != nil {
		t.Errorf("unable to add block with error %s", err.Error())
	}

	var blocks []*Block
	iter := chain.Iterator()
	block := iter.Next()

	for block != nil {
		blocks = append(blocks, block)
		block = iter.Next()
	}

	if len(blocks) != 2 {
		t.Errorf("blockchain count doesnt match. got size %d", len(blocks))
	}

	genesis := Genesis()
	if genesis.String() != blocks[1].String() {
		t.Errorf("blockchain doesn't have genesis - %s,\n got %s", genesis, blocks[1])
	}

	if string(blocks[0].Data) != "testing" {
		t.Errorf("blockchain doesn't have new block, got data %s", blocks[1].Data)
	}
}
