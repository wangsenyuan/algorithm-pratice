package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, ss [][]int, x int, expect []float64) {
	res := solve(ss, x)

	for i := 0; i < len(ss); i++ {
		if math.Abs(expect[i]-res[i])/math.Max(1, expect[i]) > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	x := 9
	ss := [][]int{
		{3, 4, 5},
		{0, 7, 8},
	}
	expect := []float64{68.538461538, 44.538461538}
	runSample(t, ss, x, expect)
}

func TestSample2(t *testing.T) {
	x := 10
	ss := [][]int{
		{1, 2, 3},
		{4, -5, 6},
		{7, -8, 9},
	}
	expect := []float64{-93.666666667, -74.666666667, -15.666666667}
	runSample(t, ss, x, expect)
}
