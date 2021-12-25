package main

import "testing"

func runSample(t *testing.T, N int, X int, A []int, edges [][]int, expect int) {
	res := solve(N, X, A, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", N, X, A, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, X := 3, 5
	A := []int{1, -5, -10}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}

	expect := -4
	runSample(t, N, X, A, edges, expect)
}