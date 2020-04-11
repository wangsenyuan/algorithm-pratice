package main

import "testing"

func runSample(t *testing.T, n uint64, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2)
	runSample(t, 2, 4)
	runSample(t, 3, 2)
	runSample(t, 4, 4)
	runSample(t, 5, 8)
}
