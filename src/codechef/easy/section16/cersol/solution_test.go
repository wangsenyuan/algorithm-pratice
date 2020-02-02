package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, edges [][]int, expect int) {
	res := solve(n, k, A, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, k, A, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	A := []int{1, 1, 2, 2, 3}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 3
	runSample(t, n, k, A, edges, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 5
	A := []int{1, 1, 2, 2, 3}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := -1
	runSample(t, n, k, A, edges, expect)
}
