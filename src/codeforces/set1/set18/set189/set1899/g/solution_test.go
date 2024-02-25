package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, p []int, queries [][]int, expect []bool) {
	res := solve(n, edges, p, queries)

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
	p := []int{1, 2, 3}
	queries := [][]int{
		{1, 2, 2},
		{1, 2, 3},
		{2, 3, 1},
		{1, 2, 3},
		{2, 3, 3},
	}
	expect := []bool{true, false, true, false, true}

	runSample(t, n, edges, p, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	edges := [][]int{
		{2, 6},
		{2, 7},
		{2, 4},
		{1, 7},
		{2, 8},
		{10, 6},
		{8, 5},
		{9, 4},
		{3, 4},
	}
	p := []int{10, 2, 5, 9, 1, 7, 6, 4, 3, 8}
	queries := [][]int{
		{8, 9, 8},
		{7, 8, 1},
		{7, 10, 6},
		{4, 8, 9},
		{5, 5, 10},
		{7, 10, 1},
		{9, 9, 2},
		{9, 10, 6},
		{6, 6, 2},
		{10, 10, 6},
	}
	expect := []bool{false, true, true, true, false, true, true, false, false, false}

	runSample(t, n, edges, p, queries, expect)
}
