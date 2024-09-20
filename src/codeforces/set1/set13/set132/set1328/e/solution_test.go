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
	n := 10
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
		{3, 7},
		{7, 8},
		{7, 9},
		{9, 10},
	}
	queries := [][]int{
		{3, 8, 9, 10},
		{2, 4, 6},
		{2, 1, 5},
		{4, 8, 2},
		{6, 10},
		{5, 4, 7},
	}
	expect := []bool{true, true, true, true, false, false}

	runSample(t, n, edges, queries, expect)
}
