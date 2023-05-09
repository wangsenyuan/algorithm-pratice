package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, pts [][]int, expect float64) {
	res := solve(pts)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %.3f, but got %.3f", expect, res)
	}

}

func TestSample1(t *testing.T) {
	pts := [][]int{
		{3, 0},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}
	expect := 2.0
	runSample(t, pts, expect)
}
