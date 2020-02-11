package main

import "testing"

func runSample(t *testing.T, n, k, m uint64, expect uint64) {
	_, res := solve(n, k, m)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, k, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 3, 500, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 2, 1000, 1)
}
