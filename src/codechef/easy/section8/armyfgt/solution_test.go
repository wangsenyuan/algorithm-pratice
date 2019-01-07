package main

import "testing"

func runSample(t *testing.T, n int, A []uint64, lower uint64, upper uint64, expect uint64) {
	res := solve(n, A, lower, upper)

	if res != expect {
		t.Errorf("Sample %d %v %d %d, expect %d, but got %d", n, A, lower, upper, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []uint64{2, 4}, 4, 8, 3)
}
