package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 6
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 7
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	edges := [][]int{
		{2, 3},
		{3, 9},
		{6, 3},
		{9, 8},
		{9, 10},
		{4, 8},
		{3, 1},
		{3, 5},
		{7, 1},
	}
	expect := 67
	runSample(t, n, edges, expect)
}
