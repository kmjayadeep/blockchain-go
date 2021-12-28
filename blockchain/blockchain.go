package blockchain

import (
	"fmt"
	"log"

	"github.com/kmjayadeep/blockchain-go/block"
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
		genesis := block.Genesis()

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

func (chain *BlockChain) AddBlock(data string) error {
	newBlock := block.CreateBlock(data, chain.LastHash)
	serialized, err := newBlock.Serialize()
	if err != nil {
		return err
	}
	err = chain.DB.Put(newBlock.HashString(), serialized)
	if err != nil {
		return err
	}
	err = chain.DB.Put("lh", newBlock.Hash)
	if err != nil {
		return err
	}
	chain.LastHash = newBlock.Hash
	return nil
}

func (chain *BlockChain) Iterator() *Iterator {
	iter := &Iterator{
		chain.DB,
		fmt.Sprintf("%x", chain.LastHash),
	}
	return iter
}
