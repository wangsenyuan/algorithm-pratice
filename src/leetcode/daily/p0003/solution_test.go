package p0003

import "testing"

func TestSample1(t *testing.T) {
	target := Constructor([]int{233})

	res := target.ShowFirstUnique()

	if res != 233 {
		t.Errorf("Sample fail at step 1, expect 233, but got %d", res)
	}

	target.Add(11)

	res = target.ShowFirstUnique()

	if res != 233 {
		t.Errorf("Sample fail at step 2, expect 233, but got %d", res)
	}
}
