package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// A transaction occurs when one or more inputs are combined to give one or more outputs.
type Transaction struct {
	ID []byte
	Vin []TXInput
	Vout []TXOutput
}


// An output of a transaction consists of a value (i.e. number of satoshis, hundred millionth of a BTC) and ScriptPubKey, which is used to "lock" this amount of satoshis.
// Important to note that outputs are indivisible, meaning that you cannot reference a part of its value. When value is referenced in a new transaction, it's spent as a whole.
// To deal with this, when the value is greater than required, a change is generated and sent back to sender.
type TXOutput struct {
	Value int
	ScriptPubKey string
}

// An input to a transaction. Inputs references previous output in bitcoin. Txid is the id of this particular transaction where is input is inputted into.
// Vout references an index of a previous output in the transaction. ScriptSig provides data to be used in an output's ScriptPubKey (i.e. to unlock the value in the output
// to be used as input to a transaction.
type TXInput struct {
	Txid []byte
	Vout int
	ScriptSig string
}


// Sets ID of a transaction, which is implemented as the hash of the entire transaction struct.
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

