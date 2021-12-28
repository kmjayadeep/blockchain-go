package main

import (
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/cli"
	"github.com/kmjayadeep/blockchain-go/storage"
)

func main() {
	// close app once this function is finished
	defer os.Exit(0)

	opts := badger.DefaultOptions("/tmp/blockchain")
	opts.Logger = nil
	db, err := storage.NewDatabase(opts)
	defer db.Close()

	if err != nil {
		log.Fatal(err, "unable to create db")
	}

	chain, err := blockchain.InitBlockChain(db)
	if err != nil {
		log.Fatal(err, "unable to initialize blockchain")
	}

	commandLine := cli.NewCommandLine(chain)
	commandLine.Run(os.Args)
}
