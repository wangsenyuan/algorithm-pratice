package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	res := solve(s)

	if math.Abs(expect-res)/math.Max(1, expect) > 1e-6 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "technocup"
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "tictictactac"
	expect := 0.333333333333333
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "bbaabaabbb"
	expect := 0.1
	runSample(t, s, expect)
}
