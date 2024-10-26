package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2, 0},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 1},
		{1, 3, 0},
		{2, 3, 1},
		{3, 4, 1},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	edges := [][]int{
		{1, 2, 0},
		{8, 3, 0},
		{2, 3, 1},
		{1, 4, 1},
		{8, 7, 0},
		{1, 5, 1},
		{4, 6, 1},
		{5, 7, 0},
		{6, 8, 0},
	}
	expect := 3
	runSample(t, n, edges, expect)
}
