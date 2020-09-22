package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	solve(n, A)

	if getMaxSum(n, A) != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, getMaxSum(n, A))
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{-4, 2, -4, 3, -5}
	var expect int64 = 5
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{-3, -2, -1}
	var expect int64 = 0
	runSample(t, n, A, expect)
}
