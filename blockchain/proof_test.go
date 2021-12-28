package blockchain

import (
	"fmt"
	"testing"
)

func TestProof(t *testing.T) {
	pow := NewProof(CreateBlock("test", []byte("test")))

	if pow.Target == nil {
		t.Errorf("target not defined")
	}

	if pow.Target.String() != "28269553036454149273332760011886696253239742350009903329945699220681916416" {
		t.Errorf("target is invalid: got : %s", pow.Target)
	}

	nonce, hash := pow.Run()

	if nonce != 13761 {
		t.Errorf("got invalid nonce")
	}

	if fmt.Sprintf("%x", hash) != "0005a106619410ca1c365c6bf02bb1a25bbcb96cd55a2dd24348b36af5d3ecd0" {
		t.Errorf("got invalid hash")
	}

}

func BenchmarkProofRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow := NewProof(CreateBlock(fmt.Sprintf("test block %d", i), []byte("prev hash")))
		pow.Run()
	}
}
