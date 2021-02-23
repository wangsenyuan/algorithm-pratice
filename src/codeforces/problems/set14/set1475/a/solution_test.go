package main

import "testing"

func runSample(t *testing.T, n uint64, expect bool) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %t, but got %t", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, false)
	runSample(t, 3, true)
	runSample(t, 4, false)
	runSample(t, 5, true)
	runSample(t, 1099511627776, false)
}
