package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect float64) {
	res := solve(n, A)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v, expect %f, but got %f", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{1, 2}
	expect := 1.5
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	expect := 4.0
	runSample(t, n, A, expect)
}
