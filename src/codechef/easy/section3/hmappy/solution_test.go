package main

import "testing"

func runSample(t *testing.T, n int, M int64, A []int, B []int, expect int64) {
	res := solve(n, M, A, B)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, epxect %d, but got %d", n, M, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 15)
}
