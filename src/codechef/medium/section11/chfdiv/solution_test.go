package main

import "testing"

func runSample(t *testing.T, N, K int64, expect int) {
	res := solve(N, K)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", N, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 6)
}
