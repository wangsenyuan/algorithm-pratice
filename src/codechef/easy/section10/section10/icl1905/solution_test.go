package main

import "testing"

func runSample(t *testing.T, N, A int, C, D []int, expect int64) {
	res := solve(N, A, C, D)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", N, A, C, D, expect, res)
	}
}

func TestSample(t *testing.T) {
	N, A := 3, 10
	C := []int{1, 2, 3}
	D := []int{3, 1, 2}
	runSample(t, N, A, C, D, 20)
}
