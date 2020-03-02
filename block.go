package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func newBlock(data string, PrevBlockHash []byte) *Block {

	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{},
	}

	block.SetHash()

	return block
}

// func main() {

// 	b := newBlock("Prueba", []byte("65df2s4"))

// 	fmt.Println("Timestamp: ", (b.Timestamp))
// 	fmt.Println("Data: ", string(b.Data))
// 	fmt.Println("PrevBlockHash: ", string(b.PrevBlockHash))
// 	fmt.Println("Hash: ", string(b.Hash))

// }
