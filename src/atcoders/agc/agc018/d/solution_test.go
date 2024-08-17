package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := int(solve(n, edges))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 5},
		{3, 4, 7},
		{2, 3, 3},
		{2, 5, 2},
	}
	expect := 38
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	edges := [][]int{
		{2, 8, 8},
		{1, 5, 1},
		{4, 8, 2},
		{2, 5, 4},
		{3, 8, 6},
		{6, 8, 9},
		{2, 7, 12},
	}
	expect := 132
	runSample(t, n, edges, expect)
}
