package main

import "testing"

func runSample(t *testing.T, n int, A [][]int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := [][]int{
		{0, 0, 2, 2},
		{1, 0, 2, 0},
		{0, 3, 0, 3},
		{2, 4, 0, 0},
	}
	var expect int64 = 13
	runSample(t, n, A, expect)
}
