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
  checkEqual(t, outputs[0].PubKey, to)
  checkEqual(t, outputs[1].PubKey, from)
  checkEqual(t, outputs[0].Value, 100)
  checkEqual(t, outputs[1].Value, 150)
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
