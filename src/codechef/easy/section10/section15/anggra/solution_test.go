package main

import "testing"

func runSample(t *testing.T, n, m, k int, expect int) {
	res := solve(n, m, k)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, m, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 1, 7)
}
