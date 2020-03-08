package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, A []int, edges [][]int, queries []string, expect []int64) {
	res := solve(N, Q, A, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 5, 8
	edges := [][]int{
		{4, 3},
		{3, 1},
		{5, 2},
		{1, 2},
	}
	A := []int{1, 10, 4, 9, 4}
	queries := []string{
		"+ 1 6",
		"? 3",
		"+ 3 5",
		"? 3",
		"+ 2 2",
		"+ 5 10",
		"? 5",
		"? 4",
	}
	expect := []int64{6, 0, 33, 25}

	runSample(t, N, Q, A, edges, queries, expect)
}
