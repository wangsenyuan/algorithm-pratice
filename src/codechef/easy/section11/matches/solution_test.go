package main

import "testing"

func runSample(t *testing.T, M, N uint64, expect bool) {
	res := solve(M, N)

	if res != expect {
		t.Errorf("Sample %d %d, expect %t, but got %t", M, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 3, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 155, 47, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 6, 4, false)
}
