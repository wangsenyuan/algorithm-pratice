package main

import "testing"

func runSample(t *testing.T, L, R uint64, K int, expect int64) {
	res := solve(L, R, K)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", L, R, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 48, 48, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 48, 48, 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 15, 1, 11)
}
