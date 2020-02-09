package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	expect := 0
	runSample(t, n, edges, expect)
}
