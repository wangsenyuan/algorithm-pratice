package main

import "testing"

func runSample(t *testing.T, N, B int64, expect int64) {
	res := solve(N, B)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", N, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 2, 6)
}
