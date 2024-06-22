package main

import "testing"

func runSample(t *testing.T, n int, k int, edges [][]int, expect int) {
	res := solve(n, k, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 1
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	expect := 2
	runSample(t, n, k, edges, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 1
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}
	expect := 3
	runSample(t, n, k, edges, expect)
}
