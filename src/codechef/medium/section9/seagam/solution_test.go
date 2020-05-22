package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, N int, A []int, a int, b float64) {
	c, d := solve(N, A)

	if a != c || math.Abs(b-d) > 1e-4 {
		t.Errorf("Sample %d %v, expect %d, %.5f, but got %d %.5f", N, A, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	A := []int{6, 10, 15, 22, 28}
	a := 0
	b := 0.4
	runSample(t, N, A, a, b)
}

func TestSample2(t *testing.T) {
	N := 5
	A := []int{2, 4, 8, 16, 32}
	a := 1
	b := 1.0
	runSample(t, N, A, a, b)
}

func TestSample3(t *testing.T) {
	N := 4
	A := []int{1, 2, 3, 4}
	a := 1
	b := 0.5833
	runSample(t, N, A, a, b)
}
