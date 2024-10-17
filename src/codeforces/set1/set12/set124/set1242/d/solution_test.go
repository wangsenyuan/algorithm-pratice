package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{2, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{3, 4},
		{3, 5},
		{3, 6},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 10

	edges := [][]int{
		{1, 5},
		{1, 8},
		{1, 9},
		{5, 8},
		{8, 9},
		{4, 7},
		{2, 3},
		{3, 10},
		{2, 6},
		{2, 10},
	}

	expect := 0
	runSample(t, n, edges, expect)
}
