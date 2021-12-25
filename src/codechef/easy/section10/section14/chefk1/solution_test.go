package main

import "testing"

func runSample(t *testing.T, n, m int64, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 6, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3, 2)
}
