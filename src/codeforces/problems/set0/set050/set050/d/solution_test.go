package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, k int, e float64, orig []int, pos [][]int, expect float64) {
	res := solve(n, k, e, orig, pos)

	if math.Abs(res-expect) > 1e-8 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	k := 1
	e := float64(500)
	orig := []int{5, 5}
	pos := [][]int{
		{1, 2},
	}
	expect := 3.84257761518762740

	runSample(t, n, k, e, orig, pos, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	k := 3
	e := float64(100)
	orig := []int{0, 0}
	pos := [][]int{
		{3, 4},
		{60, 70},
		{100, 100},
		{10, 10},
		{5, 12},
	}
	expect := 13.45126176453737600

	runSample(t, n, k, e, orig, pos, expect)
}
