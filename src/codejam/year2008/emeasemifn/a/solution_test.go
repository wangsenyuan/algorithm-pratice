package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A, B []int, x, y float64) {
	a, b := solve(A, B)

	if math.Abs(x-a) > 1e-6 || math.Abs(y-b) > 1e-6 {
		t.Errorf("Sample %v %v, expect %f, %f, but got %f, %f", A, B, x, y, a, b)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 0, 0, 10, 0, 0}
	B := []int{3, 3, 1, 1, 3, 1}
	x, y := 2.692308, 1.538462
	runSample(t, A, B, x, y)
}
