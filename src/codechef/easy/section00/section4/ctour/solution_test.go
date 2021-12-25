package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, roads [][]int, q int, queries [][]int, expect []int64) {
	res := solve(n, roads, q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, roads, q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 9, 5
	roads := [][]int{
		{1, 2, 8},
		{1, 3, -9},
		{2, 4, 1},
		{2, 5, -6},
		{3, 6, 7},
		{3, 7, 6},
		{6, 8, 3},
		{6, 9, 4},
	}
	qs := [][]int{
		{1, 2},
		{2, 7},
		{4, 3},
		{3, 2},
		{8, 9},
	}
	expect := []int64{10, 5, 0, -1, 21}

	runSample(t, n, roads, q, qs, expect)
}
