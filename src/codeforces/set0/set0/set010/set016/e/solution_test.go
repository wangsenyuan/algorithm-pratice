package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a [][]float64, expect []float64) {
	res := solve(a)
	for i := 0; i < len(a); i++ {
		if math.Abs(res[i]-expect[i])/math.Max(1, expect[i]) > 1e-6 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := [][]float64{
		{0, 0.5},
		{0.5, 0},
	}
	expect := []float64{0.500000, 0.500000}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]float64{
		{0, 1, 1, 1, 1},
		{0, 0, 0.5, 0.5, 0.5},
		{0, 0.5, 0, 0.5, 0.5},
		{0, 0.5, 0.5, 0, 0.5},
		{0, 0.5, 0.5, 0.5, 0},
	}
	expect := []float64{1.000000, 0.000000, 0.000000, 0.000000, 0.000000}
	runSample(t, a, expect)
}
