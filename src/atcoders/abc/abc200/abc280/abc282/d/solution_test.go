package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{4, 2},
		{3, 1},
		{5, 2},
		{3, 2},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 9
	edges := [][]int{
		{4, 9},
		{9, 1},
		{8, 2},
		{8, 3},
		{9, 2},
		{8, 4},
		{6, 7},
		{4, 6},
		{7, 5},
		{4, 5},
		{7, 8},
	}
	expect := 9
	runSample(t, n, edges, expect)
}
