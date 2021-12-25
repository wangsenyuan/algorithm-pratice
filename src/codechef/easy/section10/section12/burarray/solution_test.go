package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int64, q int, queries [][]int64, expect []int64) {
	res := solve(n, q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := int64(7)
	q := 9
	queries := [][]int64{
		{2, 2, 5},
		{1, 0},
		{2, -3, 0},
		{1, 1},
		{2, 0, 3},
		{1, -4},
		{1, -2},
		{2, -4, -1},
		{2, -5, -3},
	}
	expect := []int64{
		5, 4, 4, 0, 1,
	}
	runSample(t, n, q, queries, expect)
}
