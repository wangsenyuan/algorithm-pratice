package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a, b, c int, k int, expect float64) {
	res := solve(a, b, c, k)

	if math.Abs(res-expect) > 1e-8 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4, 5, 1, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 35, 75, 100, 2, 700)
}
