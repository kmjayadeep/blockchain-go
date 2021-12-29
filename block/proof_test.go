package block_test

import (
	"encoding/hex"
	"testing"

	"github.com/kmjayadeep/blockchain-go/block"
	"github.com/kmjayadeep/blockchain-go/transaction"
)

func createTestBlock(txId, prevHash string) *block.Block {
	data := []*transaction.Transaction{{
		ID: []byte(txId),
	}}
	block := block.CreateBlock(data, []byte(prevHash))
	return block
}

func TestProof(t *testing.T) {
	pow := block.NewProof(createTestBlock("testTxId", "testPrevHash"))

	if pow.Target == nil {
		t.Errorf("target not defined")
	}

	if pow.Target.String() != "28269553036454149273332760011886696253239742350009903329945699220681916416" {
		t.Errorf("target is invalid: got : %s", pow.Target)
	}

	nonce, hash := pow.Run()

	if nonce != 466 {
		t.Errorf("got invalid nonce, got %d", nonce)
	}

	hashString := hex.EncodeToString(hash)
	if hashString != "0009236aa05ed56b5d2c4c86fb14f1313f3c1cfe4fbd0772993fbdace29868b9" {
		t.Errorf("got invalid hash. got %s", hashString)
	}

	if !pow.Validate() {
		t.Errorf("pow should be valid")
	}

	pow.Block.Transactions = []*transaction.Transaction{{ID: []byte("modified")}}

	if pow.Validate() {
		t.Errorf("pow should be invalid")
	}

}

func BenchmarkProofRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow := block.NewProof(createTestBlock("testId", "testHash"))
		pow.Run()
	}
}
