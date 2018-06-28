package main

import "testing"

func runSample(t *testing.T, N int, T int, input [][]int, expect int) {
	res := solve(N, T, input)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, T, input, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, T := 7, 53
	input := [][]int{
		{5, 8, 2},
		{2, 4, 8},
		{6, 8, 13},
		{1, 13, 12},
		{4, 5, 1},
		{3, 2, 7},
		{3, 13, 5},
	}
	expect := 1
	runSample(t, N, T, input, expect)
}
