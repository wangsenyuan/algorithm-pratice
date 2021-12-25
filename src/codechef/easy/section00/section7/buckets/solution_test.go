package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, k int, A [][]int, expect []float64) {
	res := solve(n, k, A)

	for j := 0; j < k; j++ {
		if math.Abs(res[j]-expect[j]) > 1e-5 {
			t.Errorf("Sample %d %d %v, expect %v, but got %v", n, k, A, expect, res)
			return
		}
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	A := [][]int{
		{0, 1},
		{1, 1},
	}
	expect := []float64{0.333333, 0.666667}
	runSample(t, n, k, A, expect)
}
