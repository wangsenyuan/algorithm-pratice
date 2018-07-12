package main

import "testing"

func runSample(t *testing.T, A, B, P int64, expect int) {
	res := solve(A, B, P)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", A, B, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 20, 5, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 20, 3, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 365, 952, 5, 129)
}
