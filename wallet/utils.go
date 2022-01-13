package wallet

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
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

func checksum(payload []byte) []byte {

	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:checksumLength]
}
