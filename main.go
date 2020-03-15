package main

//Develop branch
import (
	"fmt"

	"github.com/NCoin/domain"
)

func main() {

	bc := domain.NewBlockchain()

	bc.AddBlock("0.6 BTC")
	bc.AddBlock("0.3 BTC")
	bc.AddBlock("0.1 BTC")

	for _, block := range bc.Blocks {

		fmt.Printf("Previous block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Actual block hash: %x\n\n", block.Hash)
	}

}
