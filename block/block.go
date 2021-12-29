package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"

	"github.com/kmjayadeep/blockchain-go/transaction"
)

type Block struct {
	Hash         []byte
	Transactions []*transaction.Transaction
	PrevHash     []byte
	Nonce        int
}

func CreateBlock(transactions []*transaction.Transaction, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		transactions,
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
	return hex.EncodeToString(b.Hash)
}

func (b *Block) hashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) String() string {
	return fmt.Sprintf("Block - Hash:%x, PrevHash:%x, Transactions:%d, Nonce:%d",
		b.Hash,
		b.PrevHash,
		len(b.Transactions),
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

func Genesis(coinbase *transaction.Transaction) *Block {
	return CreateBlock(
		[]*transaction.Transaction{coinbase},
		[]byte{})
}
