package blockchain

import (
	"log"

	"github.com/kmjayadeep/blockchain-go/storage"
)

type BlockChain struct {
	DB       storage.Database
	LastHash []byte
}

func InitBlockChain(db storage.Database) (*BlockChain, error) {

	hash, err := db.Get("lh")

	if err != nil && err != storage.ErrKeyNotFound {
		log.Println("unable to get lh", err)
		return nil, err
	}

	if err == storage.ErrKeyNotFound {
		// key doesn't exist in db
		log.Println("creating new geneis")
		genesis := Genesis()

		serialized, err := genesis.Serialize()
		if err != nil {
			return nil, err
		}

		err = db.Put(genesis.HashString(), serialized)
		if err != nil {
			return nil, err
		}

		err = db.Put("lh", genesis.Hash)
		if err != nil {
			return nil, err
		}

		hash = genesis.Hash
	}

	chain := BlockChain{
		db,
		hash,
	}
	return &chain, nil
}

func (chain *BlockChain) AddBlock(data string) {
	// prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// newBlock := CreateBlock(data, prevBlock.Hash)
	// chain.Blocks = append(chain.Blocks, newBlock)
}
