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

func TestChecksum(t *testing.T) {

	table := []struct {
		test     string
		payload  string
		expected string
	}{
		{
			test:     "test1",
			payload:  "payload",
			expected: "e78731bb",
		},
		{
			test:     "test2",
			payload:  "payload2",
			expected: "1a49dc9d",
		},
	}

	for _, tt := range table {
		t.Run(tt.test, func(t *testing.T) {
			sum := checksum([]byte(tt.payload))
			enc := hex.EncodeToString(sum)
			if enc != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, enc)
			}
		})
	}

}
