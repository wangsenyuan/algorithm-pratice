package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A[0], A[1], A[2], A[3], A[4], A[5])

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 1, 1, 2, 2, 2}, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 3, 2, 4, 2}, 1)
}
