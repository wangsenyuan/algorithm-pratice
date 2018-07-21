package main

import "testing"

func runSample(t *testing.T, N, M int, edges [][]int, expect bool) {
	res := solve(N, M, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", N, M, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 4, 4
	edges := [][]int{
		{3, 2},
		{3, 1},
		{2, 1},
		{1, 4},
	}

	expect := true

	runSample(t, N, M, edges, expect)
}
