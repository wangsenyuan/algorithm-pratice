package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect float64) {
	res := solve(n, A)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v, expect %f, but got %f", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{3, 2, 9}
	expect := 4.0
	runSample(t, n, A, expect)
}
