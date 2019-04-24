package main

import (
	"fmt"
	"github.com/yuditan/go-blockchain/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Sending 1 BTC to Yudi")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)


		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
