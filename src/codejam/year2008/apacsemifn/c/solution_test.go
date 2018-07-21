package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, p float64, x int, expect float64) {
	res := solve(n, p, x)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %f %d, expect %f, but got %f", n, p, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0.5, 500000, 0.5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 0.75, 600000, 0.843750)
}
