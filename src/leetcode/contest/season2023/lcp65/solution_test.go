package lcp65

import "testing"

func runSample(t *testing.T, ops []int, expect int) {
	res := unSuitability(ops)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := []int{5, 3, 7}
	expect := 8
	runSample(t, ops, expect)
}
