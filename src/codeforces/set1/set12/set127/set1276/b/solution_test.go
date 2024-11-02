package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, x int, y int, expect int) {
	res := solve(n, edges, x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 5},
	}
	x, y := 3, 5
	expect := 4
	runSample(t, n, edges, x, y, expect)
}
