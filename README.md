# Golang Blockchain

## Usage via [CLI](cli/cli.go):
1. Compile the code using go build, should return an executable named "go-blockchain"
2. To print the current blockchain's blocks, run "./go-blockchain printchain"
3. To add a new block to the blockchain, run "./go-blockchain addblock -data "your transaction details here"
4. For more detailed usage on the CLI, just run the executable "./go-blockchain"
5. Since boltDB is used, data in the blockchain will persist as long as the "blockchain.db" file is untouched. To
reset blockchain to a fresh chain, just delete "blockchain.db" file.


## Basic Data types
### [Block](blockchain/block.go)
- Timestamp (Unix timestamp of transaction time)
- Data (Raw data of what is being transacted)
- PreviousHash (Hash of the previous block in the blockchain)
- Nonce (A counter that starts from 0 and increments)
- Hash (Hash of the above 4 attributes)
- Has method NewBlock that will create a new block struct. 

### [Blockchain](blockchain/blockchain.go)
- Essentially a slice of blocks
- When initialized, creates a genesis block that is the first block in the slice.
- Has method AddBlock that will append a new block to the blockchain.

## Core Algorithm

### [Proof of Work](blockchain/pow.go)
- Packaged as a struct which consists of a block and the target (upper bound for which
 the block's hash is considered valid)
- Has method Run() that will run actual PoW algorithm, incrementing the nonce each iteration until
the hash of the block (described above) is less than the target.

## Database for Persistence

- Will be using go's boltDB key-value store for persistence
- We will only store 2 types on key-value pairs
    - 32-byte block hash ->Serialized Block Structure
    - 'l' -> hash of the last block in a chain
- Note that this is a much simplified version of Bitcoin Core's implementation:
    - https://en.bitcoin.it/wiki/Bitcoin_Core_0.11_(ch_2):_Data_Storage
- Will you the database as such:
    - When we call the NewBlockChain function:
        1. Open boltDB's file on disk
        2. Check if a Blockchain is stored already
        3. If Blockchain exists:
            - Create a new Blockchain instance
            - Set the tip of the Blockchain instance to the last block hash stored in DB (sort of like linked-list)
        4. If Blockchain does not exist:
            - Create genesis block
            - Store in DB
            - Save genesis block's hash as last block hash
            - Create a new Blockchain instance with tip pointing at genesis block.