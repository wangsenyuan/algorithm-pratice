package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, ships [][]int, expect float64) {
	res := solve(n, ships)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %v, expect %f, but got %f", n, ships, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	ships := [][]int{
		{0, 0, 0, 1},
		{1, 2, 0, 1},
		{3, 4, 0, 1},
		{2, 1, 0, 1},
	}
	expect := 3.5

	runSample(t, n, ships, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	ships := [][]int{{1, 1, 1, 1}}
	expect := 0.0

	runSample(t, n, ships, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	ships := [][]int{
		{1, 0, 0, 1},
		{2, 1, 1, 4},
		{3, 2, 3, 2},
	}
	expect := 2.33333333

	runSample(t, n, ships, expect)
}
