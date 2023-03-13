package main

import (
	"fmt"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
