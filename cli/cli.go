package cli

import (
	"flag"
	"fmt"
	"github.com/yuditan/go-blockchain/blockchain"
	"os"
	"strconv"
)

// Our CLI struct, only contains a pointer to the blockchain so that we can call its methods such as Iterator
type CLI struct {
	Bc *blockchain.Blockchain
}

// Instruction set for users on the usage of the CLI-- includes what flags are available and how to use them.
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA -> adds a block to the blockchain")
	fmt.Println("  printchain -> print all the blocks of the blockchain")
}


// Flag argument validation, ensures that there is at least 1 flag so that the CLI program can do its work, else we print CLI's usage instruction and exit.
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Main entry-point for our CLI, sets flags and sub-commands.
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error with addblock command: %v\n", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("Error with printchain command: %v\n", err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

// CLI's command of addblock translates to the underlying blockchain's addblock method.
func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Block added!")
}

// CLI's command of printChain translates to the underlying blockchain's iterator's Next() method.
func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	// Since we are using an iterator now, we need an infinite for loop that will loop until we reach the genesis block.
	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		// Reached genesis block
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

