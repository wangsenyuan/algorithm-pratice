package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, dh, dl, dr, k int, expect float64) {
	res := solve(dh, dl, dr, k)
	if math.Abs(res-expect) > 1e-3 {
		t.Errorf("Sample %d %d %d %d, expect %f, but got %f", dh, dl, dr, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 7, 10, 3, 17.31)
}
