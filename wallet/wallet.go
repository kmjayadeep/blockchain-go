package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

const (
	version = byte(0x0)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func MakeWallet() (*Wallet, error) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return &Wallet{
		*private,
		pub,
	}, nil
}

func (w *Wallet) Address() ([]byte, error) {
	pubHash, err := publicKeyHash(w.PublicKey)
	if err != nil {
		return nil, err
	}

	versionedHash := append([]byte{version}, pubHash...)
	csum := checksum(versionedHash)

	fullHash := append(versionedHash, csum...)

	address := base58Encode(fullHash)

	return address, nil
}

func (w *Wallet) AddressString() (string, error) {
	address, err := w.Address()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", address), nil
}
