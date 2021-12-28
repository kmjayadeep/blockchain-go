package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/storage"
)

func main() {

	db, err := storage.NewDatabase(badger.DefaultOptions("/tmp/blockchain"))

	if err != nil {
		log.Fatal(err, "unable to create db")
	}

	chain, err := blockchain.InitBlockChain(db)
	if err != nil {
		log.Fatal(err, "unable to create db")
	}

	chain.AddBlock("block1")
	chain.AddBlock("block2")
	chain.AddBlock("block2")

	iter := chain.Iterator()

	b := iter.Next()

	for b != nil {
		// fmt.Printf("Data %s\n", b.Data)
		// fmt.Printf("hash %x\n", b.Hash)
		// fmt.Printf("nonce %d\n", b.Nonce)

		fmt.Println(b.String())

		pow := block.NewProof(b)
		fmt.Printf("pow validated %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("\n\n")

		b = iter.Next()
	}

}
