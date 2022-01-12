package blockchain_test

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/storage"
	"github.com/kmjayadeep/blockchain-go/transaction"
)

var (
	GenesisData = "First transaction from Genesis"
)

func createTestTransaction(id string) *transaction.Transaction {
	return &transaction.Transaction{
		ID:      []byte(id),
		Inputs:  []transaction.TxInput{},
		Outputs: []transaction.TxOutput{},
	}
}

func TestInitBlockChain(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()
	chain, err := blockchain.InitBlockChain(db, "myaddress")

	if chain == nil {
		t.Errorf("unable to initialize chain")
	}

	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	genesisHash := "0001e82dbba3246cc5693afda356020b6263ad0932ff8344cb728941dacdbd3c"
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

	address := "address"

	chain, err := blockchain.InitBlockChain(db, address)
	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	tx := createTestTransaction("testing")
	err = chain.AddBlock([]*transaction.Transaction{tx})
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

	tx, _ = transaction.CoinbaseTx(address, GenesisData)
	genesis := block.Genesis(tx)
	if genesis.String() != blocks[1].String() {
		t.Errorf("blockchain doesn't have genesis - %s,\n got %s", genesis, blocks[1])
	}

	if blocks[0].String() != "Block - Hash:00079b63d0bc0e4ff0938996d27f729f37588a6ac44a9fcecb74e83fdd84579f, PrevHash:0004e6cfb234624214b26607c1a6a241ec43d66f1245b23467e3be27c3cad556, Transactions:1, Nonce:1492" {
		t.Errorf("blockchain doesn't have new block, got data %s", blocks[0].String())
	}
}

func TestContinueBlockChain(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()

	address := "address"
	chain, err := blockchain.InitBlockChain(db, address)

	if chain == nil {
		t.Errorf("unable to initialize chain")
	}

	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	// reuse the same db to continue blockchain
	chain, err = blockchain.ContinueBlockChain(db)

	var blocks []*block.Block
	iter := chain.Iterator()
	b := iter.Next()

	for b != nil {
		blocks = append(blocks, b)
		b = iter.Next()
	}

	if len(blocks) != 1 {
		t.Errorf("blockchain count doesnt match. got size %d", len(blocks))
	}

	tx, _ := transaction.CoinbaseTx(address, GenesisData)
	genesis := block.Genesis(tx)
	if genesis.String() != blocks[0].String() {
		t.Errorf("blockchain doesn't have genesis - %s,\n got %s", genesis, blocks[1])
	}

}

func TestContinueBlockChainError(t *testing.T) {
	db, err := storage.NewDatabase(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()

	_, err = blockchain.ContinueBlockChain(db)

	if err != blockchain.ErrNotInitialized {
		t.Errorf("Expected error :%s", blockchain.ErrAlreadyInitialized.Error())
	}
}

func TestSendWithCorrectFunds(t *testing.T) {
	db, err := storage.NewDatabase(
		badger.DefaultOptions("").WithInMemory(true).WithLogger(nil),
	)

	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()

	address := "address1"

	chain, err := blockchain.InitBlockChain(db, address)
	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	table := []struct {
		tokens      int
		expectError bool
	}{
		{
			200,
			true,
		},
		{
			101,
			true,
		},
		{
			100,
			false,
		},
		{
			99,
			false,
		},
	}

	for _, data := range table {
		err := chain.Send("address1", "address2", data.tokens)
		if data.expectError && (err == nil || err.Error() != "not enough funds") {
			t.Fatalf("expected an error, got %v, testdata : %v", err, data)
		}
		if !data.expectError && err != nil {
			t.Fatalf("didnt expect an error, got %s, testdata : %v", err.Error(), data)
		}
		if !data.expectError {
			// reverse the send
			chain.Send("address2", "address1", data.tokens)
		}
	}

}

func TestSend(t *testing.T) {
	db, err := storage.NewDatabase(
		badger.DefaultOptions("").WithInMemory(true).WithLogger(nil),
	)

	if err != nil {
		t.Errorf("unable to initialize db")
	}
	defer db.Close()

	address := "address1"

	chain, err := blockchain.InitBlockChain(db, address)
	if err != nil {
		t.Errorf("unable to initialize chain with error %s", err.Error())
	}

	initial := chain.GetBalance("address1")
	if initial != 100 {
		t.Errorf("Expected initial balance to be 100, got %v", initial)
	}

	err = chain.Send("address1", "address2", 10)
	if err != nil {
		t.Errorf("unable to send token with error %s", err.Error())
	}

	bal := chain.GetBalance("address1")
	if bal != 90 {
		t.Errorf("Expected balance to be 90, got %v", bal)
	}
	bal = chain.GetBalance("address2")
	if bal != 10 {
		t.Errorf("Expected balance to be 10, got %v", bal)
	}

	err = chain.Send("address2", "address3", 1)
	if err != nil {
		t.Errorf("unable to send token with error %s", err.Error())
	}

	bal = chain.GetBalance("address1")
	if bal != 90 {
		t.Errorf("Expected balance to be 90, got %v", bal)
	}

	bal = chain.GetBalance("address2")
	if bal != 9 {
		t.Errorf("Expected balance to be 9, got %v", bal)
	}

	bal = chain.GetBalance("address3")
	if bal != 1 {
		t.Errorf("Expected balance to be 1, got %v", bal)
	}
}
