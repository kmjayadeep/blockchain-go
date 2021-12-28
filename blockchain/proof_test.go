package blockchain

import (
	"testing"
)

func TestNewProof(t *testing.T) {
	pow := NewProof(CreateBlock("test", []byte("test")))

	if pow.Target == nil {
		t.Errorf("target not defined")
	}

	if pow.Target.String() != "28269553036454149273332760011886696253239742350009903329945699220681916416" {
		t.Errorf("target is invalid: got : %s", pow.Target)
	}

}

func TestInitData(t *testing.T) {
	pow := NewProof(CreateBlock("test", []byte("testPrev")))

	data := pow.InitData(0)

	if len(data) != 28 {
		t.Errorf("data doesn't contain all the fields, len : %d", len(data))
	}

}
