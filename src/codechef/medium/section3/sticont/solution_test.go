package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, A []int, expect int) {
	res := solve(n, edges, A)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, edges, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 4},
		{2, 3, 7},
	}
	A := []int{1, 5, 10}

	expect := 1
	runSample(t, n, edges, A, expect)
}
