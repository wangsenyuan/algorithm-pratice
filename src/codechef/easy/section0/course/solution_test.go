package main

import (
	"testing"
	"math"
)

func TestSample(t *testing.T) {
	r, R := 5, 10
	n := 3
	cones := [][]int{
		{6, 0},
		{5, 7},
		{-2, -7},
	}
	res := solve(r, R, n, cones)
	if math.Abs(res-2.720) > 0.001 {
		t.Errorf("the sample should give answer 2.720, but got %f", res)
	}
}
