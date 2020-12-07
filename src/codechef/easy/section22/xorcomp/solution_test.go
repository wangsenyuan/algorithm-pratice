package main

import "testing"

func runSample(t *testing.T, x, y, n uint64, expect uint64) {
	res := solve(x, y, n)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", x, y, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 10, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 10, 5)
}
