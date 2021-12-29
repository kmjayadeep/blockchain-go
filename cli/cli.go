package cli

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func NewCommandLine(chain *blockchain.BlockChain) *CommandLine {
	return &CommandLine{
		chain,
	}
}

func (c *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (c *CommandLine) validateArgs(args []string) {
	if len(args) < 2 {
		c.printUsage()
		runtime.Goexit()
	}
}

func (c *CommandLine) addBlock(data string) {
	// c.blockchain.AddBlock(data)
	log.Println("Added block")
}

func (c *CommandLine) printChain() {
	iter := c.blockchain.Iterator()

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

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printCnainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch args[1] {
	case "add":
		err := addBlockCmd.Parse(args[2:])
		handleErr(err)

	case "print":
		err := printCnainCmd.Parse(args[2:])
		handleErr(err)

	default:
		c.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		c.addBlock(*addBlockData)
	}

	if printCnainCmd.Parsed() {
		c.printChain()
	}

}

func handleErr(err error) {
	if err != nil {
		log.Fatalln("error : ", err.Error())
	}
}
