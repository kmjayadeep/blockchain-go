package main

import (
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmjayadeep/blockchain-go/cli"
	"github.com/kmjayadeep/blockchain-go/storage"
)

const (
	dbPath  = ".data/blocks"
)

func main() {
	// close app once this function is finished
	defer os.Exit(0)

	opts := badger.DefaultOptions(dbPath)
	opts.Logger = nil
	db, err := storage.NewDatabase(opts)
	defer db.Close()

	if err != nil {
		log.Fatal(err, "unable to create db")
	}

	commandLine := cli.NewCommandLine(db)
	commandLine.Run(os.Args)
}
