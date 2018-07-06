package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int, expect int) {
	res := solve(n, pairs)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	pairs := [][]int{
		{3, 5},
		{-1, 2},
		{2, 4},
		{-1, 3},
		{-1, 1},
	}
	runSample(t, n, pairs, 1)
}

func TestSample2(t *testing.T) {
	n := 7
	pairs := [][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{-1, 1},
		{-1, 4},
		{-1, 3},
		{-1, 2},
	}
	runSample(t, n, pairs, -1)
}
