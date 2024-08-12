package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, vs []float64, expect float64) {
	res := solve(vs[0], vs[1], vs[2], vs[3])

	if math.Abs(res-expect)/math.Max(1, expect) > eps {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	vs := []float64{0.2, 0.2, 0.6, 0.2}
	expect := 1.532000000000
	runSample(t, vs, expect)
}

func TestSample2(t *testing.T) {
	vs := []float64{0.4998, 0.4998, 0.0004, 0.1666}
	expect := 5.005050776521
	runSample(t, vs, expect)
}

func TestSample3(t *testing.T) {
	vs := []float64{0.3125, 0.6561, 0.0314, 0.2048}
	expect := 4.260163673896
	runSample(t, vs, expect)
}
