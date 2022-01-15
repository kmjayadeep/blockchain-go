package wallet

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
	"os"
)

type Store struct {
	Wallets map[string]*Wallet
}

// Initialize a new wallet store by reading the given filePath
// If the file doesn't exist, create a new file
func InitStore(filePath string) (*Store, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		ws := &Store{
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
	return Load(file)
}

func Load(r io.Reader) (*Store, error) {
	ws := &Store{}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (ws *Store) Save(w io.Writer) error {
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(w)
	return encoder.Encode(ws)
}

func (ws *Store) GetWallet(address string) *Wallet {
	return ws.Wallets[address]
}

func (ws *Store) AddWallet() (string, error) {
	w, err := MakeWallet()
	if err != nil {
		return "", err
	}

	address, err := w.AddressString()
	if err != nil {
		return "", err
	}

	ws.Wallets[address] = w
	return address, nil
}

func (ws *Store) GetAllAddresses() []string {
	var addresses []string
	for addr := range ws.Wallets {
		addresses = append(addresses, addr)
	}
	return addresses
}
