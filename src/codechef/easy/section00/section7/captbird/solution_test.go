package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, A []int, x, y int, expect float64) {
	res := solve(n, A, x, y)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v %d %d, expect %f, but got %f", n, A, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{0, 1}
	x, y := 0, 1
	expect := 0.785398163397
	runSample(t, n, A, x, y, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{0, 1}
	x, y := 100, 1
	expect := 0.000100999899
	runSample(t, n, A, x, y, expect)
}
