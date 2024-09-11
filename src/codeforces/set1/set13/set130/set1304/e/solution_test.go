package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []bool) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	queries := [][]int{
		{1, 3, 1, 2, 2},
		{1, 4, 1, 3, 2},
		{1, 4, 1, 3, 3},
		{4, 2, 3, 3, 9},
		{5, 2, 3, 3, 9},
	}
	expect := []bool{true, true, false, true, false}

	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	queries := [][]int{
		{1, 3, 1, 1, 1},
		{1, 3, 1, 1, 2},
		{1, 3, 1, 1, 3},
		{1, 3, 1, 1, 4},
		{1, 3, 1, 1, 5},
		{1, 3, 1, 1, 6},
	}
	expect := []bool{false, true, true, true, true, true}

	runSample(t, n, edges, queries, expect)
}
