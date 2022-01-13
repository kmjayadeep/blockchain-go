package wallet

import (
	"encoding/hex"
	"testing"
)

func TestPubKeyHash(t *testing.T) {

	key := []byte("testkey")
	hash, err := publicKeyHash(key)
	enc := hex.EncodeToString(hash)

	if err != nil {
		t.Errorf("got error when hashing pub, %v", err)
	}

	expected := "d502610291e4c7581f18bc85564a1d74c47b7a14"

	if enc != expected {
		t.Errorf("expected %s, got %s", expected, enc)
	}
}
