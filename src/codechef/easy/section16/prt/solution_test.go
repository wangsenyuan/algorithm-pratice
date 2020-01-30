package main

import "testing"

func runSample(t *testing.T, n int, A []int, edges [][]int, expect int) {
	res := solve(n, A, edges)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 4}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := 24
	runSample(t, n, A, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 15
	runSample(t, n, A, edges, expect)
}
