package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, m int, n int, x float64, expect float64) {
	res := solve(m, n, x)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %.10f, but got %.10f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 1, 3
	x := 1.0
	expect := 0.5403023059
	runSample(t, m, n, x, expect)
}

func TestSample2(t *testing.T) {
	m, n := 2, 3
	x := 0.0
	expect := 4.0
	runSample(t, m, n, x, expect)
}

func TestSample3(t *testing.T) {
	m, n := 30, 1
	x := 0.1
	expect := 29.8501249583
	runSample(t, m, n, x, expect)
}
