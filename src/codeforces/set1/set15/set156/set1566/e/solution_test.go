package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
		{4, 7},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	expect := 2
	runSample(t, n, edges, expect)
}
