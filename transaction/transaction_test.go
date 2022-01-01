package transaction_test

import(
  "testing"
  "reflect"

	"github.com/kmjayadeep/blockchain-go/transaction"
)

func TestTransaction(t *testing.T) {
  from := "abc"
  to := "def"
  amount := 100
  accumulated := 250
  validOutputs := map[string][]int{
    "aa": {10, 11},
  }

  tx, err := transaction.NewTransaction(from, to, amount, accumulated, validOutputs)

  if err != nil {
    t.Fatalf("transaction not created with eror %s", err.Error())
  }

  if tx.ID == nil || len(tx.ID) == 0 {
    t.Errorf("Transaction id not set")
  }

  inputs := tx.Inputs
  checkEqual(t, len(inputs), 2)
  checkEqual(t, inputs[0].Sig, from)
  checkEqual(t, inputs[1].Sig, from)
  checkDeepEqual(t, inputs[0].ID, []byte{170})
  checkDeepEqual(t, inputs[1].ID, []byte{170})
  checkEqual(t, inputs[0].Out, 10)
  checkEqual(t, inputs[1].Out, 11)

  outputs := tx.Outputs
  checkEqual(t, len(outputs), 2)
  checkEqual(t, outputs[0].PubKey, to)
  checkEqual(t, outputs[1].PubKey, from)
  checkEqual(t, outputs[0].Value, 100)
  checkEqual(t, outputs[1].Value, 150)
}

func TestTransactionAmountValidation(t *testing.T) {
  from := "abc"
  to := "def"
  amount := 100
  accumulated := 25
  validOutputs := map[string][]int{}

  _, err := transaction.NewTransaction(from, to, amount, accumulated, validOutputs)

  if err == nil {
    t.Fatalf("transaction created")
  }

  checkEqual(t, err.Error(), "not enough funds")

}


func TestCoinbase(t *testing.T) {
  tx, err := transaction.CoinbaseTx("toaddr", "testdata")
  is := tx.IsCoinbase()

  if err != nil {
    t.Fatalf("transaction not created with eror %s", err.Error())
  }
  if tx.ID == nil || len(tx.ID) == 0 {
    t.Errorf("Transaction id not set")
  }


  checkEqual(t, is, true)

  inputs := tx.Inputs
  checkEqual(t, len(inputs), 1)
  checkEqual(t, inputs[0].Out, -1)
  checkEqual(t, inputs[0].Sig, "testdata")

  outputs := tx.Outputs
  checkEqual(t, len(outputs), 1)
  checkEqual(t, outputs[0].Value, 100)
  checkEqual(t, outputs[0].PubKey, "toaddr")
}

func TestCoinbaseDefaultData(t *testing.T) {
  tx, err := transaction.CoinbaseTx("toaddr", "")

  if err != nil {
    t.Fatalf("transaction not created with eror %s", err.Error())
  }
  inputs := tx.Inputs
  checkEqual(t, inputs[0].Sig, "Coins to toaddr")
}

func TestCanUnlock(t *testing.T) {
  in := transaction.TxInput{[]byte{}, -1, "test"}

  can := in.CanUnlock("test")
  can2 := in.CanUnlock("test2")

  checkEqual(t, can, true)
  checkEqual(t, can2, false)
}

func TestCanBeUnlocked(t *testing.T) {
  out := transaction.TxOutput{100, "test"}

  can := out.CanBeUnlocked("test")
  can2 := out.CanBeUnlocked("test2")

  checkEqual(t, can, true)
  checkEqual(t, can2, false)
}

func checkDeepEqual(t *testing.T, expected, actual interface{}) {
  if !reflect.DeepEqual(expected,actual) {
    t.Errorf("expected %d, got %d", expected, actual)
  }
}

func checkEqual(t *testing.T, expected, actual interface{}) {
  if expected != actual {
    t.Errorf("expected %d, got %d", expected, actual)
  }
}
