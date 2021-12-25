package main

import "testing"

func runSample(t *testing.T, n, m, x, y int, expect int64) {
	res := solve(n, m, x, y)

	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", n, m, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 2, 2, 24)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 4, 2, 3, 94)
}
