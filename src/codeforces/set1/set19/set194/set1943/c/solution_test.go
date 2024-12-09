package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)
	if len(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, len(res))
	}
}

func TestSample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{2, 7},
		{3, 2},
		{6, 4},
		{5, 7},
		{1, 6},
		{6, 7},
	}
	expect := 3
	runSample(t, n, edges, expect)
}
