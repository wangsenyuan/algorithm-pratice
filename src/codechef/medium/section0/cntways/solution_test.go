package main

import "testing"

func runSample(t *testing.T, n, m, a, b int, expect int64) {
	res := solve(n, m, a, b)
	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", n, m, a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 1, 1, 5)
}
