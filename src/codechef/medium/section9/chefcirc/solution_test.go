package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, m int, points [][]float64, expect float64) {
	res := solve(n, m, points)

	if math.Abs(res-expect) > 1e-3 {
		t.Errorf("Sample %d %d %v, expect %.3f, but got %.3f", n, m, points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	points := [][]float64{
		{-1, -1},
		{0, 1},
		{1, -1},
	}

	expect := 1.250000

	runSample(t, n, m, points, expect)
}
