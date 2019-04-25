package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/yuditan/go-blockchain/transaction"
	"log"
	"os"
)

//Constants for boltDB
const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// Our blockchain struct is just a slice of blocks defined in block.go
type Blockchain struct {
	tip []byte
	Db *bolt.DB
}

// Mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*transaction.Transaction) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		fmt.Printf("Error occured trying to view lastHash in DB: %v\n", err)
	}

	newBlock := NewBlock(transactions, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

// We need the genesis block to initialize a new blockchain so that blocks slice is not null, previousHash is empty slice
func NewGenesisBlock(coinbase *transaction.Transaction) *Block {
	return NewBlock([]*transaction.Transaction{coinbase}, []byte{})
}

// Helper function to check if database has been initialized (i.e. whether a blockchain has been created and stored previously).
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

// NewBlockchain connects to DB and finds the blockchain, then creates the blockchain instance for this new session so that subsequent blocks can be appended to the same chain.
func NewBlockchain() *Blockchain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

// CreateBlockchain creates a new blockchain DB as well as a genesis block.
func CreateBlockchain(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := transaction.NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}


// Returns an iterator for the blockchain, where the iterator first points to the tip (i.e. newest block).
func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci := &BlockchainIterator{bc.tip, bc.Db}
	return bci
}