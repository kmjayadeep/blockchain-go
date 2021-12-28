package blockchain

import "github.com/kmjayadeep/blockchain-go/db"

type BlockChain struct {
	DB       db.Database
	LastHash []byte
}

func (chain *BlockChain) AddBlock(data string) {
	// prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// newBlock := CreateBlock(data, prevBlock.Hash)
	// chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockChain(db db.Database) *BlockChain {
	return &BlockChain{
		db,
		nil,
	}
}
