package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, c []int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, c, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	c := []int{1, 1, 1, 1, 1, 1, 1}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
		{3, 6},
		{3, 7},
	}
	queries := [][]int{
		{1, 3, 2},
		{2, 1},
		{1, 4, 3},
		{2, 1},
		{1, 2, 5},
		{2, 1},
		{1, 6, 4},
		{2, 1},
		{2, 2},
		{2, 3},
	}
	expect := []int{2, 3, 4, 5, 1, 2}

	runSample(t, n, c, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 23
	c := []int{1, 2, 2, 6, 5, 3, 2, 1, 1, 1, 2, 4, 5, 3, 4, 4, 3, 3, 3, 3, 3, 4, 6}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
		{3, 7},
		{3, 8},
		{4, 9},
		{4, 10},
		{4, 11},
		{6, 12},
		{6, 13},
		{7, 14},
		{7, 15},
		{7, 16},
		{8, 17},
		{8, 18},
		{10, 19},
		{10, 20},
		{10, 21},
		{11, 22},
		{11, 23},
	}
	queries := [][]int{
		{2, 1},
		{2, 5},
		{2, 6},
		{2, 7},
		{2, 8},
		{2, 9},
		{2, 10},
		{2, 11},
		{2, 4},
		{1, 12, 1},
		{1, 13, 1},
		{1, 14, 1},
		{1, 15, 1},
		{1, 16, 1},
		{1, 17, 1},
		{1, 18, 1},
		{1, 19, 1},
		{1, 20, 1},
		{1, 21, 1},
		{1, 22, 1},
		{1, 23, 1},
		{2, 1},
		{2, 5},
		{2, 6},
		{2, 7},
		{2, 8},
		{2, 9},
		{2, 10},
		{2, 11},
		{2, 4},
	}
	expect := []int{6, 1, 3, 3, 2, 1, 2, 3, 5, 5, 1, 2, 2, 1, 1, 1, 2, 3}

	runSample(t, n, c, edges, queries, expect)
}
