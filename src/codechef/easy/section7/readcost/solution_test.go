package main

import "testing"

func runSample(t *testing.T, n int, m int64, x int, expect int64) {
	res := solve(n, m, x)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, m, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 100, 9, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 2, 10, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, 12, 2, 11, 16)
}
