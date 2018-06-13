package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int, expect int) {
	res := solve(n, pairs)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	pairs := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
	}
	expect := 2
	runSample(t, n, pairs, expect)
}

func TestSample2(t *testing.T) {
	n := 15
	pairs := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 3},
		{1, 4},
		{2, 0},
		{2, 5},
		{2, 6},
		{3, 0},
		{3, 7},
		{3, 8},
		{4, 9},
		{5, 9},
		{6, 9},
	}
	expect := 5
	runSample(t, n, pairs, expect)
}
