package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//Constants for boltDB
const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Our blockchain struct is just a slice of blocks defined in block.go
type Blockchain struct {
	tip []byte
	db *bolt.DB
}

// Appending a block to our blockchain instance
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		fmt.Printf("Error occured trying to view DB: %v\n", err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			fmt.Printf("Error putting new block to DB: %v\n", err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			fmt.Printf("Error putting new block's hash (l) to DB: %v\n", err)
		}

		bc.tip = newBlock.Hash
		return nil
	})

	if err != nil {
		fmt.Printf("Error updating DB with new block's info: %v\n", err)
	}

}

// We need the genesis block to initialize a new blockchain so that blocks slice is not null, previousHash is empty slice
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Creating a new blockchain
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil{
		fmt.Printf("An error occured trying to open DB: %v\n", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		// Core of our function: we open our bucket that stores our blocks and if it exists,
		// we read the "l" key to get the hash of the last block
		// if bucket does not exist, we generate genesis block and a new bucket, then put genesis block into bucket by doing hash -> serialized(genesisBlock)
		// as well as set the "l" key to genesisHash
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				fmt.Printf("Error occured trying to create bucket: %v\n", err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				fmt.Printf("Error occured trying to store genesis block in DB: %v\n", err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				fmt.Printf("Error occured trying to store genesis hash in DB: %v\n", err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error occured trying to update DB: %v\n", err)
	}
	bc := Blockchain{tip, db}
	return &bc
}


// Returns an iterator for the blockchain, where the iterator first points to the tip (i.e. newest block).
func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}