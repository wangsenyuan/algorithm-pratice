package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A []int, expect float64) {
	res := solve(A)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 2.25
	runSample(t, A, expect)
}
