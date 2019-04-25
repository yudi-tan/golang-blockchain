package main

import (
	"github.com/yuditan/go-blockchain/blockchain"
	cli2 "github.com/yuditan/go-blockchain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli2.CLI{bc}

	cli.Run()

}
