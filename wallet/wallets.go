package wallet

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

// Initialize a new wallet store by reading the given filePath
// If the file doesn't exist, create a new file
func InitWallets(filePath string) (*Wallets, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		ws := &Wallets{
			Wallets: make(map[string]*Wallet),
		}
		ws.Save(file)
		return ws, nil
	}
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return LoadWallets(file)
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

func (ws *Wallets) GetWallet(address string) *Wallet {
	return ws.Wallets[address]
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string
	for addr := range ws.Wallets {
		addresses = append(addresses, addr)
	}
	return addresses
}
