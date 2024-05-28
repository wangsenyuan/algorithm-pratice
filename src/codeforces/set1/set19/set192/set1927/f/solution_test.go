package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res, _ := solve(n, edges)

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 1, 1},
		{4, 5, 1},
		{5, 6, 1},
		{6, 4, 1},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 10},
		{2, 3, 8},
		{3, 1, 5},
		{4, 5, 100},
		{5, 6, 40},
		{6, 4, 3},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2, 4},
		{5, 2, 8},
		{6, 1, 7},
		{6, 3, 10},
		{6, 5, 1},
		{3, 2, 8},
		{4, 3, 4},
		{5, 3, 6},
		{2, 6, 6},
		{5, 4, 5},
		{4, 1, 3},
		{6, 4, 5},
		{4, 2, 1},
		{3, 1, 7},
		{1, 5, 5},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{
		{2, 1, 10},
		{3, 1, 3},
		{4, 2, 6},
		{1, 4, 7},
		{2, 3, 3},
	}
	expect := 3
	runSample(t, n, edges, expect)
}
