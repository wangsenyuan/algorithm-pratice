package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, l int, d int, expect float64) {
	res := solve(n, l, d)

	if math.Abs(res-expect)/math.Max(1, expect) > 1e-7 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 2, 0.468750000000000)
}

func TestSample2(t *testing.T) {
	runSample(t, 200000, 200000, 200000, 0.999986408692793)
}
