package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, colors []int, q int, queries [][]int, expect []bool) {
	res := solve(n, edges, colors, q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v %d %v, expect %v, but got %v", n, edges, colors, q, queries, expect, res)
	}
}

func TestSample(t *testing.T) {
	n, q := 9, 4
	edges := [][]int{
		{4, 1},
		{1, 5},
		{1, 2},
		{3, 2},
		{3, 6},
		{6, 7},
		{6, 8},
		{9, 6},
	}
	colors := []int{0, 1, 0, 1, 0, 0, 1, 0, 1}
	queries := [][]int{
		{3, 2},
		{7, 3},
		{4, 0},
		{9, 5},
	}

	expect := []bool{true, true, false, false}

	runSample(t, n, edges, colors, q, queries, expect)
}
