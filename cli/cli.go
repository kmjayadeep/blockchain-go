package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/blockchain"
	"github.com/kmjayadeep/blockchain-go/storage"
	"github.com/kmjayadeep/blockchain-go/wallet"
)

const walletPath = ".data/wallet.bin"

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
	fmt.Println(" createwallet - Create a new wallet and print the address")
	fmt.Println(" listwallets - List the addresses in our wallet store")
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

	balance := chain.GetBalance(address)

	fmt.Printf("Balance of %s is : %d\n", address, balance)
}

func (c *CommandLine) createWallets() {
	store, err := wallet.InitStore(walletPath)
	if err != nil {
		log.Fatal(err, "unable to open wallet store")
	}

	addr, err := store.AddWallet()
	if err != nil {
		log.Fatal(err, "unable to add wallet")
	}

	file, err := os.OpenFile(walletPath, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err, "unable to open wallet store")
	}
	err = store.Save(file)
	if err != nil {
		log.Fatal(err, "unable to save wallet")
	}

	fmt.Printf("Wallet created with address : %s", addr)
}

func (c *CommandLine) listWallets() {
	store, err := wallet.InitStore(walletPath)
	if err != nil {
		log.Fatal(err, "unable to open wallet store")
	}

	addrs := store.GetAllAddresses()
	fmt.Printf("Addresses :\n")
	for _, addr := range addrs {
		fmt.Printf("* %s\n", addr)
	}
}

func (c *CommandLine) send(from, to string, amount int) {
	chain, err := blockchain.ContinueBlockChain(c.db)
	if err != nil {
		log.Fatal(err, "unable to continue blockchain")
	}

	err = chain.Send(from, to, amount)
	if err != nil {
		log.Fatal(err, "unable to send tokens")
	}

	fmt.Printf("success")
}

func (c *CommandLine) Run(args []string) {
	c.validateArgs(args)

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listWalletsCmd := flag.NewFlagSet("listwallets", flag.ExitOnError)

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

	case "createwallet":
		err := createWalletCmd.Parse(args[2:])
		handleErr(err)

	case "listwallets":
		err := listWalletsCmd.Parse(args[2:])
		handleErr(err)

	default:
		c.printUsage()
		runtime.Goexit()
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}

	if createWalletCmd.Parsed() {
		c.createWallets()
	}

	if listWalletsCmd.Parsed() {
		c.listWallets()
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
			sendCmd.Usage()
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
