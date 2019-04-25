package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Simplified representation of a block, only consists of the core attributes of a real blockchain block, such as merkleHash, hash of previous block and transactions/data
// In Bitcoin, Timestamp prevBlockHash and Hash are block headers, which is a separate struct from transactions (data). But we combine them as a single struct for simplicity.
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

// Creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce


	return block
}


func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatal(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Fatal(err)
	}
	return &block
}