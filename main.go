package main

import (
	"log"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/db"
)

func main() {

	db, err := db.NewDatabase(badger.DefaultOptions("/tmp/blockchain"))

	if err != nil {
		log.Fatal(err, "unable to create db")
	}

	chain := blockchain.InitBlockChain(db)

	chain.AddBlock("block1")
	chain.AddBlock("block2")
	chain.AddBlock("block2")

	// for _, block := range chain.Blocks {
	// fmt.Printf("Data %s\n", block.Data)
	// fmt.Printf("hash %x\n", block.Hash)
	// fmt.Printf("nonce %d\n", block.Nonce)

	// pow := blockchain.NewProof(block)
	// fmt.Printf("pow validated %s\n", strconv.FormatBool(pow.Validate()))

	// fmt.Printf("\n\n")
	// }

}
