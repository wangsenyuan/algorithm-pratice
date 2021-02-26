package main

import "testing"

func runSample(t *testing.T, n, k int, W [][]int, expect int) {
	res := solve(n, k, W)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	W := [][]int{
		{0, 1, 4, 2},
		{1, 0, 2, 2},
		{4, 2, 0, 3},
		{2, 2, 3, 0},
	}
	expect := 11
	runSample(t, n, k, W, expect)
}
