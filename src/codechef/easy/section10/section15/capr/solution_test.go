package main

import "testing"

func runSample(t *testing.T, N uint64, M int64, expect int64) {
	res := solve(N, M)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", N, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 60)
}
