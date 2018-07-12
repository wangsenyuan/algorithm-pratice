package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int64, Q int, queries [][]int64, expect []int64) {
	res := solve(N, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", N, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := int64(12), 6
	queries := [][]int64{
		{1, 6},
		{1, 14},
		{2, 4},
		{2, 3},
		{3, 12},
		{3, 14},
	}
	expect := []int64{
		4, 2, 2, 3, 5, 6,
	}
	runSample(t, N, Q, queries, expect)
}
