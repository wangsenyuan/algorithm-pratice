package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, expect int) {
	res := solve(n, edges, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 1
	edges := [][]int{
		{1, 2, 5},
		{3, 1, 2},
		{3, 4, 3},
	}
	expect := 8
	runSample(t, n, edges, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 7, 2
	edges := [][]int{
		{1, 2, 5},
		{1, 3, 4},
		{1, 4, 2},
		{2, 5, 1},
		{2, 6, 2},
		{4, 7, 3},
	}
	expect := 14
	runSample(t, n, edges, k, expect)
}
