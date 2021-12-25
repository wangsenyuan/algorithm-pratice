package main

import "testing"

func runSample(t *testing.T, n int, A []int64, M int, edges [][]int, expect int) {
	res := solve(n, A, M, edges)
	if res != expect {
		t.Errorf("Sample %d %v %d %v, expect %d, but got %d", n, A, M, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 4
	A := []int64{1, 2, 3, 4, 5}
	edges := [][]int{
		{2, 4},
		{1, 5},
		{1, 3},
		{3, 5},
	}

	expect := 2
	runSample(t, n, A, m, edges, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 4
	A := []int64{1, 2, 3, 4, 5}
	edges := [][]int{
		{3, 5},
		{1, 2},
		{4, 5},
		{5, 2},
	}

	expect := 3
	runSample(t, n, A, m, edges, expect)
}

