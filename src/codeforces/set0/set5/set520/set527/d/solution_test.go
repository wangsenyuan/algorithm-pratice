package main

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := solve(points)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{2, 3},
		{3, 1},
		{6, 1},
		{0, 2},
	}
	expect := 3
	runSample(t, points, expect)
}
