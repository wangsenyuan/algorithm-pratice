package p2241

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	target := Constructor()

	target.Deposit([]int{0, 0, 1, 2, 1})

	res := target.Withdraw(600)

	if !reflect.DeepEqual(res, []int{0, 0, 1, 0, 1}) {
		t.Errorf("Sample result %v, not correct", res)
	}
}
