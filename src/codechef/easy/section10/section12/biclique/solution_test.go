package main

import "testing"

func runSample(t *testing.T, n, m, k int, edges [][]int, expect bool) {
	res := solve(n, m, k, edges)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %t, but got %t", n, m, k, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 5, 10, 3
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 4},
		{3, 5},
		{4, 5},
	}
	expect := true
	runSample(t, n, m, k, edges, expect)
}
