package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
		0,
	}

	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

func (b *Block) HashString() string {
	return fmt.Sprintf("%x", b.Hash)
}

func (b *Block) String() string {
	return fmt.Sprintf("Block - Hash :%x, PrevHash:%x, Data:%s, Nonce:%d",
		b.Hash,
		b.PrevHash,
		b.Data,
		b.Nonce,
	)
}

func Deserialize(data []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
