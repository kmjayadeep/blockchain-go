package wallet

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func publicKeyHash(pubkey []byte) ([]byte, error) {

	pubHash := sha256.Sum256(pubkey)

	hasher := ripemd160.New()

	_, err := hasher.Write(pubHash[:])
	if err != nil {
		return nil, err
	}

	pubRipMd := hasher.Sum(nil)

	return pubRipMd, nil
}
