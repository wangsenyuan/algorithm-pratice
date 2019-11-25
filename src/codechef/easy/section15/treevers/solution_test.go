package main

import "testing"

func runSample(t *testing.T, n int, W []int, edges [][]int, expect int) {
	res := solve(n, W, edges)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, W, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	W := []int{1, 0, 1}
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 1
	runSample(t, n, W, edges, expect)
}
