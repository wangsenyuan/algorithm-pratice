package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, P []float64, Q [][]int, UP []float64, expect []float64) {
	res := solve(P, Q, UP)

	for i := 0; i < len(expect); i++ {
		if math.Abs(expect[i]-res[i]) > 1e-7 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	P := []float64{0.1, 0.25, 0.85, 0.43}
	Q := [][]int{
		{0, 1, 3},
		{1, 2, 4},
		{1, 1, 3},
		{1, 3, 4},
		{0, 1, 4},
		{0, 2, 3},
		{0, 3, 3},
	}
	UP := []float64{
		0, 0.73, 0.00255, 0.01258,
	}
	expect := []float64{
		0.10125000, 0.99531391, 0.99951473, 0.99998009,
	}
	runSample(t, P, Q, UP, expect)
}
