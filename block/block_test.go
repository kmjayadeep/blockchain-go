package block

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kmjayadeep/blockchain-go/transaction"
)

func createTransaction(id string) *transaction.Transaction {
	return &transaction.Transaction{
		ID:      []byte(id),
		Inputs:  []transaction.TxInput{},
		Outputs: []transaction.TxOutput{},
	}
}

func TestBlock(t *testing.T) {
	data := []*transaction.Transaction{createTransaction("testId")}
	prevHash := []byte("test")
	block := CreateBlock(data, prevHash)

	if block == nil {
		t.Fatalf("block is not created by createblock")
	}

	if !reflect.DeepEqual(data, block.Transactions) {
		t.Errorf("data not stored in the block")
	}
	if bytes.Compare(prevHash, block.PrevHash) != 0 {
		t.Errorf("prevHash not stored in block")
	}

	if block.Hash == nil || len(block.Hash) == 0 {
		t.Errorf("hash not computed")
	}

	if block.HashString() != "000421a0511726b8247a03d751d85066546018f81f5f972e1b32a1529b1bf37d" {
		t.Errorf("wrong hash, got %s", block.HashString())
	}

	if block.String() != "Block - Hash:000421a0511726b8247a03d751d85066546018f81f5f972e1b32a1529b1bf37d, PrevHash:74657374, Transactions:1, Nonce:2077" {
		t.Errorf("wrong block string, got %s", block.String())
	}
}

func TestGenesis(t *testing.T) {
	tx := createTransaction("genesis")
	block := Genesis(tx)
	if string(block.Transactions[0].ID) != "genesis" {
		t.Errorf("wrong data in genesis")
	}
	hash := block.Hash
	if hash == nil || len(hash) == 0 {
		t.Errorf("hash not computed for genesis")
	}
}

func TestSerialize(t *testing.T) {
	data := []*transaction.Transaction{createTransaction("test")}
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

func TestHashString(t *testing.T) {
	tx := createTransaction("test")
	block := CreateBlock([]*transaction.Transaction{tx}, []byte("prevHash"))

	hash := block.HashString()

	if hash != "00051166f7f109de57560c56716a0e878e9587807de0399c6481f3b92caf0e4a" {
		t.Errorf("got invalid hash string, got %s", hash)
	}

}
