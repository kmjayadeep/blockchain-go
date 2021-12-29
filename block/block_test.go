package block

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kmjayadeep/blockchain-go/transaction"
)

func TestBlock(t *testing.T) {
	data := []*transaction.Transaction{}
	prevHash := []byte("test")
	block := CreateBlock(data, prevHash)

	if block == nil {
		t.Fatalf("block is not created by createblock")
	}

	if reflect.DeepEqual(data, block.Transactions) {
		t.Errorf("data not stored in the block")
	}
	if bytes.Compare(prevHash, block.PrevHash) != 0 {
		t.Errorf("prevHash not stored in block")
	}

	if block.Hash == nil || len(block.Hash) == 0 {
		t.Errorf("hash not computed")
	}
}

func TestGenesis(t *testing.T) {
	coinbase, _ := transaction.CoinbaseTx("to", "Genesis")
	block := Genesis(coinbase)
	if string(block.Transactions[0].Inputs[0].Sig) != "Genesis" {
		t.Errorf("wrong data in genesis")
	}
	hash := block.Hash
	if hash == nil || len(hash) == 0 {
		t.Errorf("hash not computed for genesis")
	}
}

func TestSerialize(t *testing.T) {
	data := []*transaction.Transaction{}
	prevHash := []byte("testHash")
	block := CreateBlock(data, prevHash)

	serialized, err := block.Serialize()

	if err != nil {
		t.Errorf("block not serialized, %s", err.Error())
	}

	newBlock, err := Deserialize(serialized)

	if newBlock == nil || err != nil {
		t.Errorf("block not deserialized, %s", err.Error())
	}

	if reflect.DeepEqual(data, newBlock.Transactions) {
		t.Errorf("data not deserialized")
	}

	if bytes.Compare(newBlock.PrevHash, prevHash) != 0 {
		t.Errorf("prevHash not deserialized")
	}

}
