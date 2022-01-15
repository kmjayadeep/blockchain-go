package wallet_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/kmjayadeep/blockchain-go/wallet"
)

func TestInitStore(t *testing.T) {
	path, err := os.MkdirTemp("", "wallet")
	defer os.RemoveAll(path)
	if err != nil {
		t.Errorf("Unable to init wallets dir with error %v", err)
	}
	path = path + "/wallet.bin"
	ws, err := wallet.InitStore(path)
	if err != nil {
		t.Errorf("Unable to init wallets with error %v", err)
	}
	if ws == nil {
		t.Errorf("received nil wallet object")
	}
	_, err = os.Stat(path)
	if err != nil {
		t.Errorf("wallet not written to file %v", err)
	}

	w := wallet.Wallet{
		PublicKey: []byte("test"),
	}
	ws.Wallets["test"] = &w

	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		t.Errorf("unable to open file %v", err)
	}
	err = ws.Save(file)
	if err != nil {
		t.Errorf("unable to save file %v", err)
	}
	file.Close()

	loaded, err := wallet.InitStore(path)

	if loaded == nil || loaded.Wallets == nil {
		t.Fatalf("Wallets are nil")
	}

	if len(loaded.Wallets) == 0 {
		t.Errorf("Wallets are empty")
	}

	if !reflect.DeepEqual(loaded, ws) {
		t.Errorf("not stored properly. expected %v, got %v", ws, loaded)
	}
}

func TestStoreSave(t *testing.T) {
	b := bytes.Buffer{}

	ws := &wallet.Store{
		Wallets: make(map[string]*wallet.Wallet),
	}
	w := wallet.Wallet{
		PublicKey: []byte("test"),
	}
	ws.Wallets["test"] = &w

	err := ws.Save(&b)
	if err != nil {
		t.Errorf("Unable to save wallets with error %v", err)
	}

	loaded, err := wallet.Load(&b)
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
		t.Errorf("not stored properly. expected %v, got %v", ws, loaded)
	}
}

func TestGetWallet(t *testing.T) {
	wallets := map[string]*wallet.Wallet{
		"a": {},
		"b": {},
	}
	ws := wallet.Store{
		Wallets: wallets,
	}

	if !reflect.DeepEqual(ws.GetWallet("a"), wallets["a"]) {
		t.Errorf("not retrieving correct wallet")
	}

	if ws.GetWallet("c") != nil {
		t.Errorf("not returning wallet as nil")
	}
}

func TestGetAllAddresses(t *testing.T) {
	wallets := map[string]*wallet.Wallet{
		"a": {},
		"b": {},
	}
	ws := wallet.Store{
		Wallets: wallets,
	}

	addresses := ws.GetAllAddresses()
	expected := []string{"a", "b"}

	if !reflect.DeepEqual(addresses, expected) {
		t.Errorf("not returning all addrresses, got %v, expected %v", addresses, expected)
	}
}
