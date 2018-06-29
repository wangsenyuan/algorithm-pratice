package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, Q int, queries [][]int, expect []int) {
	res := solve(n, edges, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, edges, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{2, 4, 5},
		{3, 5, 10},
	}
	Q := 6
	queries := [][]int{
		{5, 4, 5},
		{5, 4, 10},
		{5, 4, 1},
		{1, 2, 1},
		{4, 1, 10},
		{1, 5, 8},
	}
	expect := []int{7, 13, 0, 1, 4, 3}

	runSample(t, n, edges, Q, queries, expect)
}
