package wallet

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
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

func base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func base58Decode(input []byte) ([]byte, error) {
	encode, err := base58.Decode(string(input[:]))
	return []byte(encode), err
}
