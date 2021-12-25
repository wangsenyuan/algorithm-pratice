package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, K []int, expect int) {
	res := solve(n, edges, K)
	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, edges, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	K := []int{1, 1, 1, 1, 1, 1, 1, 1}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{3, 6},
		{4, 7},
		{4, 8},
	}
	expect := 7
	runSample(t, n, edges, K, expect)
}
