package main

import "testing"

func runSample(t *testing.T, N, K int, expect int) {
	res := solve(N, K)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", N, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 18)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 3, 6)
}
