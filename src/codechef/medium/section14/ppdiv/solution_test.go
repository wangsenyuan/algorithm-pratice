package main

import "testing"

func runSample(t *testing.T, N uint64, expect int) {
	res := solve(N)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 14, 43)
}

func TestSample4(t *testing.T) {
	runSample(t, 23, 93)
}
