package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, k uint64, expect float64) {
	res := solve(n, k)
	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %d, expect %f, but got %f", n, k, expect, res)
	}
}

func TestSmaple1(t *testing.T) {
	runSample(t, 3, 1, 0.5)
}

func TestSmaple2(t *testing.T) {
	runSample(t, 3, 2, 1.0)
}

func TestSmaple3(t *testing.T) {
	runSample(t, 4, 1, .5)
}

func TestSmaple5(t *testing.T) {
	runSample(t, 1, 5, 1.0)
}
