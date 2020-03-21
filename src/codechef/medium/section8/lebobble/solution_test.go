package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, N, D int, B []int, P []int, expect float64) {
	res := solve(N, D, B, P)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, D := 2, 7
	B := []int{4, 7}
	P := []int{50, 50}
	expect := 0.25
	runSample(t, N, D, B, P, expect)
}

func TestSample2(t *testing.T) {
	N, D := 4, 7
	B := []int{5, 6, 1, 7}
	P := []int{25, 74, 47, 99}
	expect := 1.6049
	runSample(t, N, D, B, P, expect)
}
