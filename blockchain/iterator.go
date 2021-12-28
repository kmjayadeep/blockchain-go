package blockchain

import (
	"fmt"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/storage"
)

type Iterator struct {
	DB          storage.Database
	CurrentHash string
}

func (iter *Iterator) Next() *block.Block {
	data, err := iter.DB.Get(iter.CurrentHash)
	if err != nil {
		return nil
	}

	block, err := block.Deserialize(data)
	if err != nil {
		return nil
	}

	iter.CurrentHash = fmt.Sprintf("%x", block.PrevHash)

	return block
}
