package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample  expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	queries := [][]int{
		{1, 1, 1},
		{2, 0, 10},
		{4, 10, 100},
	}

	expect := []int{1, 11, 1, 100, 0}

	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{2, 3},
		{2, 1},
		{5, 4},
		{3, 4},
	}
	queries := [][]int{
		{2, 0, 4},
		{3, 10, 1},
		{1, 2, 3},
		{2, 3, 10},
		{1, 1, 7},
	}

	expect := []int{10, 24, 14, 11, 11}

	runSample(t, n, edges, queries, expect)
}
