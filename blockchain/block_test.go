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

  if block.Hash == nil || len(block.Hash) == 0{
    t.Errorf("hash not computed")
  }
}

func TestDeriveHash(t *testing.T) {
  data := "testData"
  prevHash := []byte("test")
  block := &Block{
    Data: []byte(data),
    PrevHash: prevHash,
  }
  block.DeriveHash()
  hash := block.Hash
  if hash == nil || len(hash) == 0{
    t.Errorf("hash not computed")
  }


  // Change data and test again
  data = "testData1"
  block.Data = []byte(data)
  block.DeriveHash()
  if bytes.Compare(hash, block.Hash) == 0 {
    t.Errorf("hash not computed based on data")
  }

  // Change prevHash and test again
  hash = block.Hash
  prevHash = []byte("test1")
  block.PrevHash = prevHash
  block.DeriveHash()
  if bytes.Compare(hash, block.Hash) == 0 {
    t.Errorf("hash not computed based on prevHash")
  }
}

func TestGenesis(t *testing.T) {
  block := Genesis()
  if string(block.Data) != "Genesis" {
    t.Errorf("wrong data in genesis %s", block.Data)
  }
  hash := block.Hash
  if hash == nil || len(hash) == 0{
    t.Errorf("hash not computed for genesis")
  }
}
