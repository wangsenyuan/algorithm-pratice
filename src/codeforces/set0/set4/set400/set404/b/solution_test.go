package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a float64, d float64, n int, expect [][]float64) {
	res := solve(a, d, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if math.Abs(res[i][j]-expect[i][j]) > 1e-6 {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 2
	var a, d float64 = 2, 5
	expect := [][]float64{
		{1.0000000000, 2.0000000000},
		{2.0000000000, 0.0000000000},
	}
	runSample(t, a, d, n, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	var a, d float64 = 4.147, 2.8819
	expect := [][]float64{
		{2.8819000000, 0.0000000000},
		{4.1470000000, 1.6168000000},
		{3.7953000000, 4.1470000000},
		{0.9134000000, 4.1470000000},
		{0.0000000000, 2.1785000000},
		{0.7034000000, 0.0000000000},
	}
	runSample(t, a, d, n, expect)
}
