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
	"github.com/kmjayadeep/blockchain-go/transaction"
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
	fmt.Println(" getbalance -address ADDRESS - Get balance for the address")
	fmt.Println(" createblockchain -address - Creates a blockchain with genesis block created for the address")
	fmt.Println(" print - Prints the blocks in the chain")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT - send token")
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

func (c *CommandLine) getBalance(address string) {
	chain, err := blockchain.ContinueBlockChain(c.db)
	if err != nil {
		log.Fatal(err, "unable to continue blockchain")
	}

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s is : %d\n", address, balance)
}

func (c *CommandLine) send(from, to string, amount int) {
	chain, err := blockchain.ContinueBlockChain(c.db)
	if err != nil {
		log.Fatal(err, "unable to continue blockchain")
	}

	acc, outputs := chain.FindSpendableOutputs(from, amount)
	tx, _ := transaction.NewTransaction(from, to, amount, acc, outputs)
	chain.AddBlock([]*transaction.Transaction{tx})

	fmt.Printf("success")
}

func (c *CommandLine) Run(args []string) {
	c.validateArgs(args)

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to create blockchain with")
	sendFrom := sendCmd.String("from", "", "The address from")
	sendTo := sendCmd.String("to", "", "The address to")
	sendAmount := sendCmd.Int("amount", 0, "The amount to send")

	switch args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(args[2:])
		handleErr(err)

	case "print":
		err := printChainCmd.Parse(args[2:])
		handleErr(err)

	case "send":
		err := sendCmd.Parse(args[2:])
		handleErr(err)

	case "getbalance":
		err := getBalanceCmd.Parse(args[2:])
		handleErr(err)

	default:
		c.printUsage()
		runtime.Goexit()
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		c.createBlockchain(*createBlockchainAddress)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		c.getBalance(*getBalanceAddress)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount == 0 {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		c.send(*sendFrom, *sendTo, *sendAmount)
	}

}

func handleErr(err error) {
	if err != nil {
		log.Fatalln("error : ", err.Error())
	}
}
