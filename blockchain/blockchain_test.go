package blockchain

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/storage"
	"github.com/kmjayadeep/blockchain-go/transaction"
)

func TestInitBlockChain(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()
	chain, err := InitBlockChain(db, "myaddress")

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

	var blocks []*block.Block
	iter := chain.Iterator()
	b := iter.Next()

	for b != nil {
		blocks = append(blocks, b)
		b = iter.Next()
	}

	if len(blocks) != 1 {
		t.Errorf("blockchain doesn't have genesis, got size %d", len(blocks))
	}

	if !blocks[0].Transactions[0].IsCoinbase() {
		t.Errorf("blockchain doesn't have genesis coinbase")
	}

}

func TestAddBlock(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()
	chain, err := InitBlockChain(db, "address")
	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	err = chain.AddBlock("testing")
	if err != nil {
		t.Errorf("unable to add block with error %s", err.Error())
	}

	var blocks []*block.Block
	iter := chain.Iterator()
	b := iter.Next()

	for b != nil {
		blocks = append(blocks, b)
		b = iter.Next()
	}

	if len(blocks) != 2 {
		t.Errorf("blockchain count doesnt match. got size %d", len(blocks))
	}

	tx, _ := transaction.CoinbaseTx("address", "genesis")
	genesis := block.Genesis(tx)
	if genesis.String() != blocks[1].String() {
		t.Errorf("blockchain doesn't have genesis - %s,\n got %s", genesis, blocks[1])
	}

	if blocks[0].Transactions[0].Inputs[0].Sig != "testing" {
		t.Errorf("blockchain doesn't have new block, got data %s", blocks[0].Transactions[0].Inputs[0].Sig)
	}
}
