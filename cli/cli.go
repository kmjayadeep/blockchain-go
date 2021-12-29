package cli

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/storage"
)

type CommandLine struct {
	db *storage.DB
}

func NewCommandLine(db *storage.DB) *CommandLine {
	return &CommandLine{
		db,
	}
}

func (c *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" getbalance -address ADDRESS - get balance for the address")
	fmt.Println(" createblockchain -address creates a blockchain with genesis block created for the address")
	fmt.Println(" print - Prints the blocks in the chain")
	fmt.Println(" send -from FROM -to TO -amount AMOUNt - send token")
}

func (c *CommandLine) validateArgs(args []string) {
	if len(args) < 2 {
		c.printUsage()
		runtime.Goexit()
	}
}

func (c *CommandLine) createBlockchain(address string) {
	_, err := blockchain.InitBlockChain(c.db, address)
	if err != nil {
		fmt.Printf("Unable to create blockchain due to error : %s", err.Error())
		return
	}
	fmt.Printf("Blockchain created")
}

func (c *CommandLine) printChain() {
	chain, err := blockchain.ContinueBlockChain(c.db)
	if err != nil {
		log.Fatal(err, "unable to continue blockchain")
	}
	iter := chain.Iterator()

	b := iter.Next()

	for b != nil {
		fmt.Println(b.String())

		pow := block.NewProof(b)
		fmt.Printf("pow validated %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("\n\n")

		b = iter.Next()
	}
}

func (c *CommandLine) Run(args []string) {
	c.validateArgs(args)

	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(args[2:])
		handleErr(err)

	case "print":
		err := printChainCmd.Parse(args[2:])
		handleErr(err)

	default:
		c.printUsage()
		runtime.Goexit()
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}

	if createBlockchainCmd.Parsed() {
		c.createBlockchain("koo")
	}

}

func handleErr(err error) {
	if err != nil {
		log.Fatalln("error : ", err.Error())
	}
}
