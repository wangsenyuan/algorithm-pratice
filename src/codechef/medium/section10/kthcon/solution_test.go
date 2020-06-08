package main

import "testing"

func runSample(t *testing.T, n int, points [][]int, expect int64) {
	res := solve(n, points)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	points := [][]int{
		{2, 2},
		{-2, -2},
		{2, -2},
		{-2, 2},
		{0, 1},
	}
	var expect int64 = 28
	runSample(t, n, points, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	points := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
	}
	var expect int64 = -1
	runSample(t, n, points, expect)
}
