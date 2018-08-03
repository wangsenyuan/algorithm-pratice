package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, C, N int, expect float64) {
	res := solve(C, N)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d, expect %f, but got %f", C, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 3.0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 2.5)
}
