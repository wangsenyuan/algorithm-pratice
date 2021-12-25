package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, expect float64) {
	res := solve(n)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d, expect %f, but got %f", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	expect := 1.33333333
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	expect := 2.0
	runSample(t, n, expect)
}
