package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, ss []string, expect float64) {
	res := solve(n, ss)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %v, expect %f, but got %f", n, ss, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	ss := []string{
		"aba",
		"abc",
		"bab",
		"aac",
	}
	expect := 1.1250000
	runSample(t, n, ss, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	ss := []string{
		"aba",
		"baab",
		"abc",
	}
	expect := 0.0277778
	runSample(t, n, ss, expect)
}
