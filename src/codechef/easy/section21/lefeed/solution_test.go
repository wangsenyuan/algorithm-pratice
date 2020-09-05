package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, m int, A []int, expect float64) {
	res := solve(n, m, A)

	if math.Abs(res-expect) > 1e-9 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	A := []int{0, 0}
	expect := 1.5
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 7
	A := []int{0, 1, 1, 0}
	expect := 2.285714286
	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 2
	A := []int{0}
	expect := 1.0
	runSample(t, n, m, A, expect)
}
