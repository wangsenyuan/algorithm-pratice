package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, s []string, expect float64) {
	targets := make([]target, len(s))
	for i, x := range s {
		targets[i] = parse(x)
	}

	res := solve(targets)

	if math.Abs(res-expect)/math.Max(1, expect) > 1e-6 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"2 5 1 0.955925",
		"9 -9 14 0.299977",
		"0 1 97 0.114582",
		"-4 -2 66 0.561033",
		"0 -10 75 0.135937",
	}
	expect := 1.7674770000
	runSample(t, s, expect)
}
