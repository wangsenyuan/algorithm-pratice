package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, A, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 2, 1, 1, 3, 3}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{4, 6},
	}
	queries := [][]int{
		{3, 2, 1},
		{1, 4, 2},
		{3, 2, 2},
		{2, 6},
		{3, 2, 2},
		{1, 1, 4},
		{3, 2, 1},
	}
	expect := []int{2, 3, 1, 1}
	runSample(t, n, A, edges, queries, expect)
}
