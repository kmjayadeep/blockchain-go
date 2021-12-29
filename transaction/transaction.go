package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (tx *Transaction) SetID() error {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	if err != nil {
		return err
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

	return nil
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 &&
		len(tx.Inputs[0].ID) == 0 &&
		tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount, accumulated int, validOutputs map[string][]int) (*Transaction, error) {
	var inputs []TxInput
	var outputs []TxOutput

	if accumulated < amount {
		return nil, fmt.Errorf("not enough funds")
	}

	for txid, outs := range validOutputs {
		txId, err := hex.DecodeString(txid)
		if err != nil {
			return nil, err
		}

		for _, out := range outs {
			input := TxInput{txId, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})
	if accumulated > amount {
		outputs = append(outputs, TxOutput{accumulated - amount, from})
	}

	tx := &Transaction{nil, inputs, outputs}
	tx.SetID()
	return tx, nil
}

func CoinbaseTx(to, data string) (*Transaction, error) {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	err := tx.SetID()
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
