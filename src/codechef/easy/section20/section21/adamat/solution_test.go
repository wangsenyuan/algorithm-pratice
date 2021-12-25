package main

import "testing"

func runSample(t *testing.T, n int, mat [][]int, expect int) {
	res := solve(n, mat)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", mat, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	mat := [][]int{
		{1, 2, 9, 13},
		{5, 6, 10, 14},
		{3, 7, 11, 15},
		{4, 8, 12, 16},
	}
	expect := 2
	runSample(t, n, mat, expect)
}
