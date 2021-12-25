package main

import "testing"

func runSample(t *testing.T, n int, k int, edges [][]int, expect int) {
	res := solve(n, k, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 2
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := 6
	runSample(t, n, k, edges, expect)
}
