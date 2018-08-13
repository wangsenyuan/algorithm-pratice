package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, m int, points [][]int, expect float64) {
	res := solve(n, m, points)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d %v, expect %f, but got %f", n, m, points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 1
	points := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	expect := 5.656854249492380
	runSample(t, n, m, points, expect)
}
