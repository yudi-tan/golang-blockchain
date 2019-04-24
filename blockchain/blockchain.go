package blockchain

// Our blockchain struct is just a slice of blocks defined in block.go
type Blockchain struct {
	Blocks []*Block
}

// Appending a block to our blockchain instance
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// We need the genesis block to initialize a new blockchain so that blocks slice is not null, previousHash is empty slice
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Creating a new blockchain
func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{[]*Block{genesisBlock}}
}
