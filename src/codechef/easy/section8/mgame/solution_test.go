package main

import "testing"

func runSample(t *testing.T, N, P int, expect int64) {
	res := solve(N, P)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", N, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 13)
}
