package ch1

import "testing"

func TestAdapter(t *testing.T) {
	iprocess := Adapter{}
	t.Logf("Calling the adapter traget using the Iprocess interface.\n")
	iprocess.process()

}
