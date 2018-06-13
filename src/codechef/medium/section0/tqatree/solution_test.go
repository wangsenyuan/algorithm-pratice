package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, edges, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 2},
		{2, 3, 4},
		{4, 2, 3},
		{5, 4, 1},
	}
	queries := [][]int{
		{2, 5, 3},
		{1, 3, 1},
		{2, 5, 3},
	}
	expect := []int{8, 6}
	runSample(t, n, edges, queries, expect)
}
