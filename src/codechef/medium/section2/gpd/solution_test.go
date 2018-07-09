package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, R uint, RK uint, edges [][]uint, Q int, queries [][]uint, expect [][]uint) {
	res := solve(N, R, RK, edges, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %d %v, expect %v, but got %v", N, R, RK, edges, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 6, 4
	R, RK := 1, 2
	edges := [][]uint{
		{5, 1, 3},
		{2, 1, 4},
		{3, 2, 5},
		{4, 2, 1},
		{6, 3, 3},
	}
	queries := [][]uint{
		{1, 4, 2},
		{6, 0, 12, 0},
		{7, 12, 7},
		{4, 0, 7},
	}
	expect := [][]uint{
		{0, 6},
		{2, 7},
		{0, 1},
	}
	runSample(t, N, uint(R), uint(RK), edges, Q, queries, expect)
}
