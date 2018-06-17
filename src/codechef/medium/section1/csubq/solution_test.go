package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q, L, R int, queries [][]int, expect []uint64) {
	res := solve(N, Q, L, R, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d %v, expect %v, but got %v", N, Q, L, R, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q, L, R := 5, 6, 1, 10
	queries := [][]int{
		{1, 1, 2},
		{2, 1, 5},
		{1, 3, 11},
		{1, 4, 3},
		{2, 3, 5},
		{2, 1, 5},
	}
	expect := []uint64{5, 2, 4}

	runSample(t, N, Q, L, R, queries, expect)
}
