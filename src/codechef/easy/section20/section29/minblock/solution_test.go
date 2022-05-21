package main

import "testing"

func runSample(t *testing.T, n int, k int, E [][]int, expect int) {
	res := solve(n, k, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 8, 4
	E := [][]int{
		{1, 2, 0},
		{1, 3, 1},
		{1, 6, 1},
		{3, 4, 0},
		{3, 5, 1},
		{4, 7, 1},
		{4, 8, 1},
	}
	expect := 1
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 3
	E := [][]int{
		{1, 2, 0},
		{1, 3, 0},
		{2, 4, 0},
		{2, 5, 1},
		{3, 6, 1},
	}
	expect := -1
	runSample(t, n, k, E, expect)
}
