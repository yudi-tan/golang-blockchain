package transaction

import (
	"fmt"
)

// Subsidy is the amount of reward for mining a new block. In BTC this number is calculate dynamically (i.e. total # blocks / 21000), but we store as constant for now.
const subsidy = 10000

// In bitcoin inputs are produced from outputs, but inputs reference previous outputs. This is a classic chicken or egg problem; however, in Bitcoin, the egg (output)
// comes first. Some outputs can be produced without any inputs (such as the coins that are generated when a new block is mined, i.e. think coin minting).
// A coinbase transaction is a special type of transaction that creates outputs (coins) out of nowhere (no inputs required).
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward for mining a block given to: '%s'", to)
	}

	// In a coinbase transaction, our implementation sets Txid to empty, Vout to -1 (since it isn't referencing any previous outputs).
	// Coinbase transaction also doesnt store a script in ScriptSig, so we just store any arbitrary data.
	txin :=  TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}