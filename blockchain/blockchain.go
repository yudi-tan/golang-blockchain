package blockchain

// Our blockchain struct is just a slice of blocks defined in block.go
type Blockchain struct {
	blocks []*Block
}

// Appending a block to our blockchain instance
