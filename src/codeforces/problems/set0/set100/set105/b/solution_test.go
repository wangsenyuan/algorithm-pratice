package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, k int, A int, S [][]int, expect float64) {
	res := solve(n, k, A, S)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, A := 5, 6, 100
	S := [][]int{
		{11, 80},
		{14, 90},
		{23, 70},
		{80, 30},
		{153, 70},
	}
	expect := 1.0
	runSample(t, n, k, A, S, expect)
}

func TestSample2(t *testing.T) {
	n, k, A := 5, 3, 100
	S := [][]int{
		{11, 80},
		{14, 90},
		{23, 70},
		{80, 30},
		{153, 70},
	}
	expect := 0.9628442962
	runSample(t, n, k, A, S, expect)
}

func TestSample3(t *testing.T) {
	n, k, A := 1, 3, 20
	S := [][]int{
		{20, 20},
	}
	expect := 0.7500000000
	runSample(t, n, k, A, S, expect)
}
