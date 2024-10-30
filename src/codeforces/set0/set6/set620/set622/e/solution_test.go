package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 12
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
		{3, 7},
		{3, 8},
		{3, 9},
		{8, 10},
		{8, 11},
		{8, 12},
	}
	expect := 6
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	expect := 1
	runSample(t, n, edges, expect)
}
