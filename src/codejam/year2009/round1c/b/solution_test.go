package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, flies [][]int, expectD float64, expectT float64) {
	d, dt := solve(n, flies)
	if math.Abs(d-expectD) > 1e-6 || math.Abs(dt-expectT) > 1e-6 {
		t.Errorf("Sample %d %v, expect %f %f, but got %f %f", n, flies, expectD, expectT, d, dt)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	flies := [][]int{
		{3, 0, -4, 0, 0, 3},
		{-3, -2, -1, 3, 0, 0},
		{-3, -1, 2, 0, 3, 0},
	}
	expectD := 0.00000000
	expectT := 1.0
	runSample(t, n, flies, expectD, expectT)
}

func TestSample2(t *testing.T) {
	n := 3
	flies := [][]int{
		{-5, 0, 0, 1, 0, 0},
		{-7, 0, 0, 1, 0, 0},
		{-6, 3, 0, 1, 0, 0},
	}
	expectD := 1.00000000
	expectT := 6.0
	runSample(t, n, flies, expectD, expectT)
}

func TestSample3(t *testing.T) {
	n := 4
	flies := [][]int{
		{1, 2, 3, 1, 2, 3},
		{3, 2, 1, 3, 2, 1},
		{1, 0, 0, 0, 0, -1},
		{0, 10, 0, 0, -10, -1},
	}
	expectD := 3.36340601
	expectT := 1.0
	runSample(t, n, flies, expectD, expectT)
}
