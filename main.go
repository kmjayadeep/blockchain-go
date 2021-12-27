package main

import (
	"fmt"
	"strconv"

	"github.com/kmjayadeep/blockchain-go/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock("block1")
	chain.AddBlock("block2")
	chain.AddBlock("block2")

	for _, block := range chain.Blocks {
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("hash %x\n", block.Hash)
		fmt.Printf("nonce %d\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("pow validated %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("\n\n")
	}

}
