package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, points [][]int, expect float64) {
	res := solve(points)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	expect := 2.828427

	runSample(t, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{1, 4},
		{2, 3},
		{5, 2},
		{3, 5},
		{4, 1},
	}
	expect := 9.640986

	runSample(t, points, expect)
}
