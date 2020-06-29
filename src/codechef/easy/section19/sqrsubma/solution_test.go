package main

import "testing"

func runSample(t *testing.T, n int, A []int, X int, expect int64) {
	res := solve(n, A, X)

	if res != expect {
		t.Errorf("Sample %d %v %d, expect %d, but got %d", n, A, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 1, 12}
	X := 36
	expect := int64(6)
	runSample(t, n, A, X, expect)
}
