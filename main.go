package main

import (
	"fmt"

	"github.com/kmjayadeep/blockchain-go/blockchain"
)

func main() {

	chain := blockchain.InitBLockChain()

	chain.AddBlock("block1")
	chain.AddBlock("block2")
	chain.AddBlock("block2")

	for _, block := range chain.Blocks {
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("hash %x\n", block.Hash)
		fmt.Printf("\n\n")
	}

}
