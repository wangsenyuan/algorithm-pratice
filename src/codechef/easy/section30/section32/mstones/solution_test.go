package main

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := solve(points)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{1, 1},
		{3, 1},
	}
	expect := 3
	runSample(t, points, expect)
}
