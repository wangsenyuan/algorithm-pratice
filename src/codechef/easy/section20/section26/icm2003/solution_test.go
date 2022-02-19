package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, b, c float64, expect float64) {
	res := solve(b, c)
	if math.Abs(res-expect) > 10e-9 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 5.8831725615)
}
