package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, m int, ps [][]float64, expect float64) {
	res := solve(n, m, ps)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d %v, expect %f, but got %f", n, m, ps, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 3
	ps := [][]float64{
		{0.5000, 0.5000, 0.5000},
		{0.5000, 0.5000, 0.5000},
	}
	expect := float64(0.5)
	runSample(t, n, m, ps, expect)
}
