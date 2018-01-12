package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect float64) {
	res := solve(n, edges)
	if math.Abs(expect-res) > 0.001 {
		t.Errorf("sample %d %v, should give %f, but got %f", n, edges, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 3},
		{1, 3, 2},
	}
	expect := 2.333333
	runSample(t, n, edges, expect)
}
