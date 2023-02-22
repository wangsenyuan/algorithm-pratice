package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, P []float64, expect float64) {
	res := solve(P)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []float64{0.5, 0.5, 0.5}
	expect := 2.75
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []float64{0.7, 0.2, 0.1, 0.9}
	expect := 2.489200000000000
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []float64{1, 1, 1, 1, 1}
	expect := 25.0
	runSample(t, P, expect)
}
