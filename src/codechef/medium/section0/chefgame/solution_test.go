package main

import "testing"

func runSample(t *testing.T, N, D int, points [][]int, expect int64) {
	res := solve(N, D, points)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, D, points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, D := 3, 2
	points := [][]int{
		{0, 0},
		{-1, -1},
		{1, -1},
	}

	runSample(t, N, D, points, 8)
}

func TestSample2(t *testing.T) {
	N, D := 2, 3
	points := [][]int{
		{0, 0, 0},
		{19265, 19189, 2849},
	}
	runSample(t, N, D, points, 0)
}
