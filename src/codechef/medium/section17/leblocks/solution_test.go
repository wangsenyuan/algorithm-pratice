package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A, C []int, k int, expect float64) {
	res := solve(A, C, k)

	if math.Abs(res-expect) > 1e-5 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 2, 4, 1, 1, 4}
	C := []int{1, 2, 1, 3, 2, 3, 4}
	k := 5
	expect := 0.9095238

	runSample(t, A, C, k, expect)
}
