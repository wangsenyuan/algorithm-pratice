package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	queries := [][]int{
		{3, 1, 2},
		{3, 1, 3},
		{3, 2, 3},
		{2, 2},
		{3, 1, 2},
		{3, 1, 3},
		{3, 2, 3},
	}
	expect := []int{1, 2, 1, 1, -1, -1}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 5},
		{6, 4},
		{2, 3},
		{3, 5},
		{5, 6},
	}
	queries := [][]int{
		{3, 3, 4},
		{2, 5},
		{3, 2, 6},
		{3, 1, 2},
		{2, 3},
		{3, 3, 1},
	}
	expect := []int{3, -1, 3, 2}
	runSample(t, n, edges, queries, expect)
}
