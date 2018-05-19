package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int64) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 1, 32)
}
