package main

import (
	"app/blockchain"
)

func main() {
	
	blockchain := blockchain.New()
	blockchain.AddBlock("Block 1")
	blockchain.AddBlock("Block 2")
	blockchain.AddBlock("Block 3")
	blockchain.AddBlock("Block 4")

	blockchain.PrintBlocks()

}