package main

import "testing"

func runSample(t *testing.T, N int, A []int, expect int64) {
	res := solve(N, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{1, 1, 1, 1}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, []int{1, 4}, 4)
}
