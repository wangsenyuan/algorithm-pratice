package main

import "testing"

func runSample(t *testing.T, N int, roads [][]int, F []int, P int, Q int, expect int64) {
	res := solve(N, roads, F, P, Q)

	if res != expect {
		t.Errorf("Sample %d %v %v %d %d, expect %d, but got %d", N, roads, F, P, Q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{2, 4, 7},
	}
	F := []int{3, 6, 2, 2}
	P, Q := 1, 4
	expect := int64(28)
	runSample(t, n, roads, F, P, Q, expect)
}
