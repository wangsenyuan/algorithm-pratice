package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
	runSample(t, 2, 1)
	runSample(t, 3, 500000005)
	runSample(t, 4, 944444453)
	runSample(t, 5, 616319451)
}
