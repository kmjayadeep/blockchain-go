package blockchain

import (
	"reflect"
	"testing"
)

func TesInitBlockChain(t *testing.T) {
	chain := InitBlockChain()

	if len(chain.Blocks) != 1 {
		t.Errorf("chain should only have Genesis block initially")
	}

	genesis := Genesis()

	if !reflect.DeepEqual(genesis, chain.Blocks[0]) {
		t.Errorf("fist block should be genesis")
	}

}
