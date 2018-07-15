package main

import "testing"

func runSample(t *testing.T, N, M, A int, can bool) {
	res := solve(N, M, A)
	resCan := res != nil

	if resCan != can {
		t.Errorf("Sample %d %d %d, expect %t, but got %t", N, M, A, can, resCan)
	}
}

func TestSample1(t *testing.T) {
	N, M, A := 1, 1, 1
	runSample(t, N, M, A, true)
}

func TestSample2(t *testing.T) {
	N, M, A := 1, 2, 64
	runSample(t, N, M, A, false)
}

func TestSample3(t *testing.T) {
	N, M, A := 10, 10, 1
	runSample(t, N, M, A, true)
}
