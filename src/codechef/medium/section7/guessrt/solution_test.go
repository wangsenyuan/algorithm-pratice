package main

import "testing"

func runSample(t *testing.T, N, K, M int64, expect int64) {
	res := solve(N, K, M)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", N, K, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 9, 1, 400000003)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 9, 2, 196428573)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 20, 3, 555555560)
}
