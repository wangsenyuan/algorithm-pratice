package main

import "testing"

func runSample(t *testing.T, M int, A []int, B []int, N int, expect int64) {
	res := solve(M, A, B, N)
	if res != expect {
		t.Errorf("Sample %d %v %v %d, expect %d, but got %d", M, A, B, N, expect, res)
	}
}

func TestSample(t *testing.T) {
	M, N := 3, 3
	A := []int{1, 2, 3}
	B := []int{4, 5, 6}
	runSample(t, M, A, B, N, 63)
}
