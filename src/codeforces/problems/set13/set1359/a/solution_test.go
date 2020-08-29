package main

import "testing"

func runSample(t *testing.T, n, m, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, m, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 3, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 4, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 6, 3, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 42, 0, 7, 0)
}
