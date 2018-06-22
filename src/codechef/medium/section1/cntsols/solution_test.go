package main

import "testing"

func runSample(t *testing.T, U, D, M, N int, expect int64) {
	res := solve(U, D, M, N)

	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", U, D, M, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 3, 5, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2013, 3, 31, 1)
}
