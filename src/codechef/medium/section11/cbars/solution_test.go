package main

import "testing"

func runSample(t *testing.T, n int, m uint64, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 14)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3, 322)
}
