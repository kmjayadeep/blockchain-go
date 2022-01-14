package wallet_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kmjayadeep/blockchain-go/wallet"
)

func TestWalletsStore(t *testing.T) {
	b := bytes.Buffer{}

	ws := wallet.NewWallets()
	w := wallet.Wallet{
		PublicKey: []byte("test"),
	}
	ws.Wallets["test"] = &w

	err := ws.Save(&b)
	if err != nil {
		t.Errorf("Unable to save wallets with error %v", err)
	}

	loaded, err := wallet.LoadWallets(&b)
	if err != nil {
		t.Errorf("Unable to load wallets with error %v", err)
	}

	if loaded == nil || loaded.Wallets == nil {
		t.Fatalf("Wallets are nil")
	}

	if len(loaded.Wallets) == 0 {
		t.Errorf("Wallets are empty")
	}

	if !reflect.DeepEqual(loaded, ws) {
		t.Errorf("not stored properly")
	}
}
