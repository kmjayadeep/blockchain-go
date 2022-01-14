package wallet

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() *Wallets {
	return &Wallets{
		Wallets: make(map[string]*Wallet),
	}
}

func LoadWallets(r io.Reader) (*Wallets, error) {
	ws := &Wallets{}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (ws *Wallets) Save(w io.Writer) error {
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(w)
	return encoder.Encode(ws)
}
