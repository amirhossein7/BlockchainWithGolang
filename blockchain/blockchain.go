package blockchain

import (
	"app/blockchain/block"
	"fmt"
)

type blockChain struct {
	blocks []*block.Block
}

func (bc blockChain) GetBlocks() []*block.Block {
	return bc.blocks
}
func (chain *blockChain) AddBlock(data string) {
	perviousHash := chain.blocks[len(chain.blocks)-1].GetHash()
	newBlock := block.New(data, perviousHash)
	chain.blocks = append(chain.blocks, newBlock)
}
func (chain *blockChain) PrintBlocks() {
	for _, v := range chain.GetBlocks() {
		fmt.Printf("data: %s \n", v.GetData())
		fmt.Printf("hash: %x \n", v.GetHash())
		fmt.Printf("previous hash: %x \n", v.GetPrevHash())
	}
}

func New() *blockChain {
	return &blockChain{[]*block.Block{block.Genesis()}}
}
