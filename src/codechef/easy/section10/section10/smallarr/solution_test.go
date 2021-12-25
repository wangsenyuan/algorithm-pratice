package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, x int, A []int, expect float64) {
	res := solve(n, x, A)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d %v, expect %f, but got %f", n, x, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 3, 2
	A := []int{1, -2, 3}
	runSample(t, n, x, A, 0.5)
}
