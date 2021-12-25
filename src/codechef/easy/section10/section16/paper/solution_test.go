package main

import (
	"math"
	"testing"
)

func runClosePairDistance(t *testing.T, n int, points [][]int, expect float64) {
	res := closePairDist(n, points)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v, expect %f, but got %f", n, points, expect, res)
	}
}

func TestClosePairDist(t *testing.T) {
	n := 6
	points := [][]int{
		{2, 3}, {12, 30}, {40, 50}, {5, 1}, {12, 10}, {3, 4},
	}
	expect := 1.414214
	runClosePairDistance(t, n, points, expect)
}

func runSample(t *testing.T, n int, m int, w int, h int, S string, points [][]int, expect float64) {
	res := solve(n, m, w, h, S, points)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, w, h := 6, 2, 10, 10
	S := "ULRDDL"
	points := [][]int{
		{4, 4},
		{5, 5},
	}
	expect := 1.41421356237
	runSample(t, n, m, w, h, S, points, expect)
}

func TestSample2(t *testing.T) {
	n, m, w, h := 4, 2, 10, 10
	S := "RUDL"
	points := [][]int{
		{1, 1},
		{9, 9},
	}
	expect := 2.00000000000
	runSample(t, n, m, w, h, S, points, expect)
}
