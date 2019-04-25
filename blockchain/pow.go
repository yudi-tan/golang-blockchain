package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/yuditan/go-blockchain/utils"
	"math"
	"math/big"
)

// Mining Difficulty, represents the number of leading 0's in the final hash of a block, changes from time to time.
// For simplicity sake, won't implement target adjusting algorithm, only use a constant.
const targetBits = 22
const maxNonce = math.MaxInt64

// Proof Of Work struct encapsulating a block and the target it must achieve to be considered valid.
type ProofOfWork struct {
	block *Block
	target *big.Int
}

// Initialize a new proof
// 256 is length of sha-256 hash in bits
// By shifting our target left by targetBits, we can then check for every subsequent block hash if the blockhash is less than
// this target (shifting left by targetBits results in a large number). We only allow block hashes which are smaller than this target to be considered valid.
// Lowering the target (i.e. an upper bound on validity of block hash) will result in fewer valid numbers, and thus increasing the difficulty of PoW.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

// Helper function to combine the various block attributes with the none into a single byte slice so that it can be hashed later on
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			utils.IntToHex(int64(pow.block.Timestamp)),
			utils.IntToHex(int64(nonce)),
	},
	[]byte{},
	)

	return data
}

// Actual PoW Algorithm
// Initializes nonce to 0 and repeatedly checks if combined hash is less than target. If greater than target, increment nonce and recompare.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

// PoW Validation
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}