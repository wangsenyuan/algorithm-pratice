package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, M, Q int, qs [][]float64, expect float64) {
	res := solve(M, Q, qs)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d %v, expect %f, but got %f", M, Q, qs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	M, Q := 10, 2
	qs := [][]float64{
		{0.25, 0.25, 0.25, 0.25},
		{0.25, 0.25, 0.25, 0.25},
	}
	expect := 0.625
	runSample(t, M, Q, qs, expect)
}

func TestSample2(t *testing.T) {
	M, Q := 64, 3
	qs := [][]float64{
		{0.3, 0.4, 0.0, 0.3},
		{1.0, 0.0, 0.0, 0.0},
		{0.2, 0.2, 0.2, 0.4},
	}
	expect := 1.0
	runSample(t, M, Q, qs, expect)
}

func TestSample3(t *testing.T) {
	M, Q := 3, 2
	qs := [][]float64{
		{0.5, 0.17, 0.17, 0.16},
		{0.5, 0.25, 0.25, 0.0},
	}
	expect := 0.5
	runSample(t, M, Q, qs, expect)
}
