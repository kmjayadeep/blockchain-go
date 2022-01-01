package blockchain

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/storage"
	"github.com/kmjayadeep/blockchain-go/transaction"
)

type BlockChain struct {
	DB       storage.Database
	LastHash []byte
}

var (
	GenesisData           = "First transaction from Genesis"
	ErrNotInitialized     = fmt.Errorf("Blockchain not initialized")
	ErrAlreadyInitialized = fmt.Errorf("Blockchain already initialized")
)

// Continue an existing blockchain from db
func ContinueBlockChain(db storage.Database) (*BlockChain, error) {

	hash, err := db.Get("lh")

	if err != nil && err != storage.ErrKeyNotFound {
		log.Println("unable to get lh", err)
		return nil, err
	}

	if err == storage.ErrKeyNotFound {
		return nil, ErrNotInitialized
	}

	chain := BlockChain{
		db,
		hash,
	}
	return &chain, nil
}

// Initialize a new blockchain by adding a genesis block
func InitBlockChain(db storage.Database, address string) (*BlockChain, error) {
	hash, err := db.Get("lh")

	if err != nil && err != storage.ErrKeyNotFound {
		log.Println("unable to get lh", err)
		return nil, err
	}
	if err != storage.ErrKeyNotFound {
		return nil, ErrAlreadyInitialized
	}

	// key doesn't exist in db
	log.Println("creating new genesis")
	coinbase, err := transaction.CoinbaseTx(address, GenesisData)
	if err != nil {
		return nil, err
	}
	genesis := block.Genesis(coinbase)

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

	chain := BlockChain{
		db,
		hash,
	}
	return &chain, nil
}

// Find transactions which have unspent tokens for a given address
func (chain *BlockChain) findUnspentTransactions(address string) []transaction.Transaction {
	var unspentTxs []transaction.Transaction

	spentTXOs := make(map[string][]int)
	iter := chain.Iterator()

	for {
		block := iter.Next()

		if block == nil {
			break
		}

		for _, tx := range block.Transactions {
			txId := hex.EncodeToString(tx.ID)

			for outIdx, out := range tx.Outputs {
				spent := false
				if spentTXOs[txId] != nil {
					for _, spentOut := range spentTXOs[txId] {
						if spentOut == outIdx {
							spent = true
							break
						}
					}
				}
				if !spent && out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}

			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Out)
					}
				}
			}

		}
	}

	return unspentTxs
}

// Find unspent transaction outpits for a given address
func (chain *BlockChain) findUTXO(address string) []transaction.TxOutput {
	var UTXOs []transaction.TxOutput

	unspentTransactions := chain.findUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// Find spendable outputs for a given address
func (chain *BlockChain) findUsableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := chain.findUnspentTransactions(address)

	accumulated := 0

	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Outputs {
			if out.CanBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)
			}
			if accumulated >= amount {
				return accumulated, unspentOuts
			}
		}
	}

	return accumulated, unspentOuts
}

func (chain *BlockChain) GetBalance(address string) int {
	balance := 0
	UTXOs := chain.findUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}

func (chain *BlockChain) Send(from, to string, amount int) error {
	acc, outputs := chain.findUsableOutputs(from, amount)
	tx, err := transaction.NewTransaction(from, to, amount, acc, outputs)
	if err != nil {
		return err
	}
	return chain.AddBlock([]*transaction.Transaction{tx})
}

// Add a new block to blockchain with given set of transactions
func (chain *BlockChain) AddBlock(transactions []*transaction.Transaction) error {
	newBlock := block.CreateBlock(transactions, chain.LastHash)
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
