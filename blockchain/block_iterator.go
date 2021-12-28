package blockchain

import (
	"fmt"

	"github.com/kmjayadeep/blockchain-go/storage"
)

type BlockIterator struct {
	DB          storage.Database
	CurrentHash string
}

func (iter *BlockIterator) Next() *Block {
	data, err := iter.DB.Get(iter.CurrentHash)
	if err != nil {
		return nil
	}

	block, err := Deserialize(data)
	if err != nil {
		return nil
	}

	iter.CurrentHash = fmt.Sprintf("%x", block.PrevHash)

	return block
}
