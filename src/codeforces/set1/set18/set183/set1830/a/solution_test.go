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
		{4, 5},
		{1, 3},
		{1, 2},
		{3, 4},
		{1, 6},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{5, 6},
		{2, 4},
		{2, 7},
		{1, 3},
		{1, 2},
		{4, 5},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	edges := [][]int{
		{2, 1},
		{4, 2},
		{7, 3},
		{2, 6},
		{2, 5},
		{1, 3},
	}
	expect := 2
	runSample(t, n, edges, expect)
}
