package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/yuditan/go-blockchain/transaction"
	"log"
	"time"
)

// Simplified representation of a block, only consists of the core attributes of a real blockchain block, such as merkleHash, hash of previous block and transactions/data
// In Bitcoin, Timestamp prevBlockHash and Hash are block headers, which is a separate struct from transactions (data). But we combine them as a single struct for simplicity.
type Block struct {
	Timestamp int64
	Transactions []*transaction.Transaction
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

// Creates a new block
func NewBlock(transactions []*transaction.Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce


	return block
}

// Hash transactions in a block into a single hash by concatenating all hashes of each transaction (note each transaction's ID is the entire transaction's hash)
// and then hashing the concatenated combination. In bitcoin, this is implemented use Merkle Trees instead, which allows for quick checking
// whether a block contains certain transactions, and having only the root hash without downloading all the transactions.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}


// Serialize block struct into byte slice
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatal(err)
	}
	return result.Bytes()
}

// Deserialize byte slice into block struct
func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Fatal(err)
	}
	return &block
}