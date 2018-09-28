package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, points [][]int, expect float64) {
	res := solve(n, points)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v, expect %.7f, but got %.7f", n, points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = []int{i, i}
	}
	expect := 5.656854
	runSample(t, n, points, expect)
}
