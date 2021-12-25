package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, X []int, V []int, expect float64) {
	res := solve(n, X, V)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v %v, expect %f, but got %f", n, X, V, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	X := []int{10, 20, 30}
	V := []int{10, 30, 10}
	var expect float64 = 0.5
	runSample(t, n, X, V, expect)
}
