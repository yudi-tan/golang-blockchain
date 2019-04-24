package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Simplified representation of a block, only consists of the core attributes of a real blockchain block, such as merkleHash, hash of previous block and transactions/data
// In Bitcoin, Timestamp prevBlockHash and Hash are block headers, which is a separate struct from transactions (data). But we combine them as a single struct for simplicity.
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

// SHA-256 hashing of Timestamp, Data and PrevBlockHash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

