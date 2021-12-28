package blockchain

import (
	"bytes"
	"testing"
)

func TestBlock(t *testing.T) {
	data := "testData"
	prevHash := []byte("test")
	block := CreateBlock(data, prevHash)

	if block == nil {
		t.Fatalf("block is not created by createblock")
	}

	if bytes.Compare([]byte(data), block.Data) != 0 {
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
	block := Genesis()
	if string(block.Data) != "Genesis" {
		t.Errorf("wrong data in genesis %s", block.Data)
	}
	hash := block.Hash
	if hash == nil || len(hash) == 0 {
		t.Errorf("hash not computed for genesis")
	}
}
