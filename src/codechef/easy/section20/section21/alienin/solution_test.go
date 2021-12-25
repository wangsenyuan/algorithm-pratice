package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, D int, C []int, expect float64) {
	res := solve(n, D, C)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	D := 2
	C := []int{3, 2, 3}
	var expect = 1.5
	runSample(t, n, D, C, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	D := 1
	C := []int{5, 6}
	var expect = 2.0
	runSample(t, n, D, C, expect)
}
