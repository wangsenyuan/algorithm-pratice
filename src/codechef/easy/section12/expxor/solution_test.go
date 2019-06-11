package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, b []int, p []float64, expect float64) {
	res := solve(n, b, p)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v %v, expect %f, but got %f", n, b, p, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	b := []int{5, 6, 2}
	p := []float64{1.0, 0.0, 0.5}
	expect := 6.0
	runSample(t, n, b, p, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	b := []int{2, 2, 2, 2}
	p := []float64{0.5, 0.5, 0.5, 0.5}
	expect := 1.0
	runSample(t, n, b, p, expect)
}
