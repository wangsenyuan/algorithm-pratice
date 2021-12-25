package main

import "testing"

func runSample(t *testing.T, L, R uint64, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 8, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 10, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 100, 2738)
}
