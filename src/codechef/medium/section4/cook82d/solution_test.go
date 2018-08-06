package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, M int, edges [][]int, Q int, queries [][]int, expect []bool) {
	res := solve(N, M, edges, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %d %v, expect %v, but got %v", N, M, edges, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 4, 6
	edges := [][]int{
		{1, 2},
		{3, 1},
		{2, 3},
		{4, 1},
		{2, 4},
		{1, 2},
	}
	Q := 3
	queries := [][]int{
		{1, 3},
		{2, 4},
		{1, 6},
	}
	expect := []bool{true, false, true}

	runSample(t, N, M, edges, Q, queries, expect)
}
