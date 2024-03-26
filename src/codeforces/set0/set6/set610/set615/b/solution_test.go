package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	edges := [][]int{
		{4, 5},
		{3, 5},
		{2, 5},
		{1, 2},
		{2, 8},
		{6, 7},
	}
	expect := 9
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := 12
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 3},
		{2, 4},
		{4, 5},
		{5, 3},
		{2, 1},
		{1, 4},
		{3, 2},
	}
	expect := 9
	runSample(t, n, edges, expect)
}
