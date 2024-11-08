package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
		{3, 7},
		{3, 8},
		{4, 9},
		{4, 10},
	}
	expect := 8
	runSample(t, n, edges, expect)
}
