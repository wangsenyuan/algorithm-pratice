package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, pairs [][]int, expect int) {
	res := solve(n, edges, pairs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	pairs := [][]int{
		{1, 3},
		{4, 1},
	}
	expect := 1
	runSample(t, n, edges, pairs, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{3, 1},
		{6, 1},
		{5, 2},
		{4, 2},
	}
	pairs := [][]int{
		{3, 1},
		{3, 6},
		{2, 6},
	}
	expect := 2
	runSample(t, n, edges, pairs, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 4
	runSample(t, n, edges, pairs, expect)
}
