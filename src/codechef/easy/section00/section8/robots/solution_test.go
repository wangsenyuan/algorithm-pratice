package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, S string, Q int, queries [][]int, expect [][]float64) {
	res := solve(S, Q, queries)

	for i := 0; i < Q; i++ {
		a, b := expect[i][0], expect[i][1]
		c, d := res[i][0], res[i][1]

		if math.Abs(a-c) > 1e-6 || math.Abs(b-d) > 1e-6 {
			t.Fatalf("Sample %s %d %v, expect %v, but got %v", S, Q, queries, expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	S := "012"
	Q := 3
	queries := [][]int{
		{1, 1},
		{1, 2},
		{2, 2},
	}
	expect := [][]float64{
		{1.00000000, 0.00000000},
		{1.50000000, 0.86602540},
		{0.50000000, 0.86602540},
	}
	runSample(t, S, Q, queries, expect)
}
