package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, P []int, A []int, B []int, expect float64) {
	res := solve(P, A, B)

	diff := math.Abs(res - expect)
	if diff/math.Max(1, expect) > 1e-6 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{3, 3}
	A := []int{1, 0}
	B := []int{-1, 6}
	expect := 3.6055512755
	runSample(t, P, A, B, expect)
}

func TestSample2(t *testing.T) {
	P := []int{3, 3}
	A := []int{-1, -1}
	B := []int{4, 3}
	expect := 3.2015621187
	runSample(t, P, A, B, expect)
}
