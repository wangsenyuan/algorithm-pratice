package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []bool) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
	}
	queries := [][]int{
		{3, 2, 5},
		{1, 2, 3, 4, 5},
		{1, 4},
		{1, 3, 5},
		{1, 5, 4},
	}
	expect := []bool{true, false, true, false, true}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{3, 2},
		{2, 4},
		{5, 2},
	}
	queries := [][]int{
		{3, 1},
		{3, 4, 5},
		{2, 3, 5},
		{1},
	}
	expect := []bool{true, false, true, true}
	runSample(t, n, edges, queries, expect)
}
