package main

import "testing"

func runSample(t *testing.T, X, Y, L, R uint64, expect uint64) {
	res := solve(X, Y, L, R)

	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", X, Y, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 12, 4, 17, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 12, 0, 8, 7)
}
