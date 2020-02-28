package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, H, W int, grid [][]int, Q int, queries [][]int, expect []int) {
	res := solve(H, W, grid, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %d %v, expect %v, but got %v", H, W, grid, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	H, W, Q := 5, 5, 3
	grid := [][]int{
		{4, 3, 9, 7, 2},
		{8, 6, 5, 2, 8},
		{1, 7, 3, 4, 3},
		{2, 2, 4, 5, 6},
		{9, 9, 9, 9, 9},
	}
	queries := [][]int{
		{3, 4, 6},
		{3, 2, 5},
		{1, 4, 9},
	}
	expect := []int{10, 0, 19}
	runSample(t, H, W, grid, Q, queries, expect)
}

func TestSample2(t *testing.T) {
	H, W, Q := 5, 5, 4
	grid := [][]int{
		{4, 3, 9, 7, 2},
		{8, 6, 5, 2, 8},
		{1, 7, 3, 4, 3},
		{2, 2, 4, 5, 6},
		{9, 9, 9, 9, 9},
	}
	queries := [][]int{
		{3, 4, 6},
		{3, 2, 5},
		{1, 4, 9},
		{5, 5, 10},
	}
	expect := []int{10, 0, 19, 25}
	runSample(t, H, W, grid, Q, queries, expect)
}
