package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
)

// Since we are now using boltDB and our blockchain is no longer a slice of blocks,
// in order to print out our chain in the order they take in the blockchain, we need to load the info from the DB.
// However, we don't want to load the entire DB into memory since that might be too large, so we create an iterator.
type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

// Iterator will return the next block from a blockchain, one block per call to Next() (which is really the previous block in the blockchain since we point from newest block)
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
	})

	if err != nil {
		fmt.Printf("Error finding newest block from DB: %v\n", err)
	}

	i.currentHash = block.PrevBlockHash
	return block
}