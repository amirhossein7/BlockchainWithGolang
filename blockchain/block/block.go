package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	data     string
	hash     []byte
	prevHash []byte
}

func (b Block) GetHash() []byte {
	return b.hash
}
func (b Block) GetData() string {
	return b.data
}
func (b Block) GetPrevHash() []byte {
	return b.prevHash
}

func (block *Block) deriveHash() {
	info := bytes.Join([][]byte{block.hash, block.prevHash}, []byte{})
	hash := sha256.Sum256(info)
	block.hash = hash[:]
}

func New(data string, prevHash []byte) *Block {
	block := &Block{data, []byte{}, prevHash}
	block.deriveHash()
	return block
}

func Genesis() *Block {
	return New("Genesis", []byte{})
}
