# Golang Blockchain

## Basic Data types
### Block
- Timestamp (Unix timestamp of transaction time)
- Data (Raw data of what is being transacted)
- PreviousHash (Hash of the previous block in the blockchain)
- Nonce (A counter that starts from 0 and increments)
- Hash (Hash of the above 4 attributes)
- Has method NewBlock that will create a new block struct. 

### Blockchain
- Essentially a slice of blocks
- When initialized, creates a genesis block that is the first block in the slice.
- Has method AddBlock that will append a new block to the blockchain.

## Core Algorithm

### Proof of Work
- Packaged as a struct which consists of a block and the target (upper bound for which
 the block's hash is considered valid)
- Has method Run() that will run actual PoW algorithm, incrementing the nonce each iteration until
the hash of the block (described above) is less than the target.