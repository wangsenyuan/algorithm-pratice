package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, vals [][]int, expect int) {
	res := solve(n, edges, vals)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	vals := [][]int{
		{1, 6},
		{3, 8},
	}
	edges := [][]int{
		{1, 2},
	}
	expect := 7
	runSample(t, n, edges, vals, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	vals := [][]int{
		{1, 3},
		{4, 6},
		{7, 9},
	}
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 8
	runSample(t, n, edges, vals, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	vals := [][]int{
		{3, 14},
		{12, 20},
		{12, 19},
		{2, 12},
		{10, 17},
		{3, 17},
	}
	edges := [][]int{
		{3, 2},
		{6, 5},
		{1, 5},
		{2, 6},
		{4, 6},
	}
	expect := 62
	runSample(t, n, edges, vals, expect)
}
