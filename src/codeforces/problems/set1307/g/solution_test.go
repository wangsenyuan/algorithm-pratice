package main

import (
	"math"
	"testing"
)

func TestSample1(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{1, 2, 2},
		{2, 3, 2},
		{1, 3, 3},
	}
	x := []int{0, 1, 2, 3, 4}
	expect := []float64{3, 4, 4.5, 5, 5.5}

	solver := NewSolver(n, m, edges)

	for i := 0; i < len(x); i++ {
		r := solver.Query(x[i])

		if math.Abs(r-expect[i]) > 1e-6 {
			t.Errorf("Sample fail at %d, expect %f, but got %f", i, expect[i], r)
		}
	}
}
