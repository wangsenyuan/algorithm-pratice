package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, r, b, p int, expect float64) {
	res := solve(r, b, p)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d %d, expect %f, but got %f", r, b, p, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 2, 2.0)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 6, 7, 1.0000000000)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 4, 2.2000000000)
}

func TestSample4(t *testing.T) {
	runSample(t, 15, 14, 13, 14.4482758621)
}

func TestSample5(t *testing.T) {
	runSample(t, 100000, 100000, 0, 100000.0000000000)
}
