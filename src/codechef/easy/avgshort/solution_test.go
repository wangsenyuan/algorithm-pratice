package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, m int, edges [][]int, s, e int, expect float64) {
	res := solve(n, m, edges, s, e)
	if math.Abs(res-expect) > 1e-8 {
		t.Errorf("sample %d %d %v %d %d, expect %f, but got %f", n, m, edges, s, e, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{3, 2, 3},
	}
	s, e := 1, 3
	runSample(t, n, m, edges, s, e, 1.5)
}
