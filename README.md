#Golang Blockchain

## Basic Data types
### Block
- Timestamp (Unix timestamp of transaction time)
- Data (Raw data of what is being transacted)
- PreviousHash (Hash of the previous block in the blockchain)
- Hash (Hash of the above 3 attributes)
- Has method NewBlock that will create a new block struct. 

### Blockchain
- Essentially a slice of blocks
- When initialized, creates a genesis block that is the first block in the slice.
- Has method AddBlock that will append a new block to the blockchain.