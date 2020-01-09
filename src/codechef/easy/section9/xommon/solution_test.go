package main

import "testing"

func runSample(t *testing.T, n int, A []uint64, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	A := []uint64{1, 200, 3, 0, 400, 4, 1, 7}
	expect := 6
	runSample(t, n, A, expect)
}
