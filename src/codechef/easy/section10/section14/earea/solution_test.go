package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, P [][]int, expect float64) {
	res := solve(n, P)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
	}
	expect := 0.125
	runSample(t, n, P, expect)
}
