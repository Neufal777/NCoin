package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)

}

func NewGenesisBlock() *Block {

	return newBlock("Bloque genesis", []byte{})
}

func NewBlockchain() *Blockchain {

	return &Blockchain{[]*Block{NewGenesisBlock()}}
}