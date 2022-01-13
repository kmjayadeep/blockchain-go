package wallet_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"reflect"
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

	key := w.PrivateKey.PublicKey
	publicKey := append(key.X.Bytes(), key.Y.Bytes()...)

	if !reflect.DeepEqual(publicKey, w.PublicKey) {
		t.Errorf("invalid publickey, got %v, expected %v", w.PublicKey, publicKey)
	}
}

func TestWalletAddress(t *testing.T) {
	pubKey, _ := hex.DecodeString("7381724bcc3ca2757ff18a17bf19de955d7ec670978296b843242954cd4df0ce79e68aa61bf7898f0ccfbb72ee928d73c4ced17eb7becad8144ef00c17a186b7")

	w := wallet.Wallet{
		ecdsa.PrivateKey{},
		pubKey,
	}

	address, err := w.Address()

	if err != nil {
		t.Errorf("got error when getting address, %v", err)
	}

	expected := "18vVoVqR3DuFL3Eh8eFWTjV5iyfCEueCVB"

	if string(address) != expected {
		t.Errorf("invalid wallet address, expected %v, got %s", expected, address)
	}
}
