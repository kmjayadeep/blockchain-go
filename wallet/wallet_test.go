package wallet_test

import (
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
