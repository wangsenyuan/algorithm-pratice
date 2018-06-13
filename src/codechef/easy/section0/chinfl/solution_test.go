package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n, m int, buy, sell [][]int, D int64, expect float64) {
	res := solve(n, m, buy, sell, D)

	if math.Abs(res-expect) > 0.00000005 {
		t.Errorf("sample %d %d %v %v %d, expect %.7f, but got %.7f", n, m, buy, sell, D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	sell := [][]int{
		{2, 2, 10},
		{5, 4, 8},
		{7, 6, 6},
	}
	buy := [][]int{
		{1, 1, 9},
		{3, 3, 7},
		{6, 5, 5},
	}

	runSample(t, n, m, buy, sell, 5, 15.0)
}
