package main

import "testing"

func runSample(t *testing.T, n int, points [][]int) {
	res := solve(n, points)

	if len(res) != n/10 {
		t.Errorf("Sample %d %v, expect %d points, but got %v", n, points, n/10, res)
	}
}

func TestSample1(t *testing.T) {
	n := 11
	points := [][]int{
		{0, 0},
		{1, 1},
		{2, 3},
		{2, 5},
		{0, 10},
		{-2, 10},
		{-5, 9},
		{-8, 7},
		{-8, 4},
		{-6, 1},
		{-2, 0},
	}
	runSample(t, n, points)
}
