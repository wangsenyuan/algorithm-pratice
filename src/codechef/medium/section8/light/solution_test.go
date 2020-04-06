package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, N int, L, R float64, light [][]float64, expect float64) {
	res := solve(N, L, R, light)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, L, R := 2, 3.2, 7.3
	light := [][]float64{
		{3.2, 4.7, 28},
		{7.3, 4.2, 75},
	}
	expect := 3.300759642
	runSample(t, N, L, R, light, expect)
}
