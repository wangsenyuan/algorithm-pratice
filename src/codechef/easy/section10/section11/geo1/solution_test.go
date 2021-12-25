package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, P [][]int, Q [][]int, expect []float64) {
	res := solve(P, Q)

	for i := 0; i < len(expect); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-7 {
			t.Errorf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	P := [][]int{
		{1, 1},
		{2, 1},
		{2, 2},
		{1, 2},
	}
	Q := [][]int{
		{1, 1},
	}
	expect := []float64{9}
	runSample(t, P, Q, expect)
}

func TestSample2(t *testing.T) {
	P := [][]int{
		{1, 1},
		{2, 1},
		{1, 2},
	}
	Q := [][]int{
		{1, 1},
		{2, 3},
	}
	expect := []float64{9.7426406, 230.8086578}
	runSample(t, P, Q, expect)
}
