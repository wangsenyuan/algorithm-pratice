package main

import "testing"

func runSample(t *testing.T, N int, T int, M int, buses [][]int, expect int) {
	res := solve(N, T, M, buses)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, T, M, buses, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, T, M := 5, 10, 5
	buses := [][]int{
		{1, 2, 1, 2},
		{1, 5, 3, 4},
		{2, 4, 4, 5},
		{2, 5, 5, 6},
		{4, 5, 6, 7},
	}
	expect := 2
	runSample(t, N, T, M, buses, expect)
}
