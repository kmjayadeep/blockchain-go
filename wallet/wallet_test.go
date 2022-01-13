package wallet_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"

	"github.com/kmjayadeep/blockchain-go/wallet"
)

func TestMakeWallet(t *testing.T) {

	w, err := wallet.MakeWallet()

	if err != nil {
		t.Errorf("got error when creating wallet, %v", err)
	}

	if w.PublicKey == nil || len(w.PublicKey) == 0 {
		t.Errorf("public key should be defined")
	}

}

func TestWalletAddress(t *testing.T) {
	pubKey := "testkey"

	w := wallet.Wallet{
		ecdsa.PrivateKey{},
		[]byte(pubKey),
	}

	address, err := w.Address()

	if err != nil {
		t.Errorf("got error when getting address, %v", err)
	}

	encAddr := hex.EncodeToString(address)
	expected := "314c52486e397032346a385739337971556f6f48644c31325747324a766a6d417a50"

	if encAddr != expected {
		t.Errorf("invalid wallet address, expected %v, got %v", expected, encAddr)
	}
}
