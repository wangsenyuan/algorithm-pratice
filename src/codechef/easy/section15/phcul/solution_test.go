package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, A []int, m int, B []int, k int, C []int, x, y int, expect float64) {
	res := solve(n, A, m, B, k, C, x, y)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %v %d %v %d %v %d %d, expect %.7f, but got %.7f", n, A, m, B, k, C, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 1, 4
	n, m, k := 3, 2, 2
	A := []int{4, 4, 2, 0, 8, 1}
	B := []int{10, 1, 3, 1}
	C := []int{1, 3, 9, 5}
	expect := float64(8.1820424980)
	runSample(t, n, A, m, B, k, C, x, y, expect)
}

func TestSample2(t *testing.T) {
	x, y := 6, 4
	n, m, k := 2, 2, 3
	A := []int{7, 10, 5, 7}
	B := []int{1, 6, 2, 3}
	C := []int{1, 8, 0, 7, 0, 2}
	expect := float64(8.6995968482)
	runSample(t, n, A, m, B, k, C, x, y, expect)
}
