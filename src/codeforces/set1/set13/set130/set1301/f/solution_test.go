package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, grid [][]int, queries [][]int, expect []int16) {
	res := solve(k, grid, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	grid := [][]int{
		{1, 2, 1, 3},
		{4, 4, 5, 5},
		{1, 2, 1, 3},
	}
	queries := [][]int{
		{1, 1, 3, 4},
		{2, 2, 2, 2},
	}
	expect := []int16{2, 0}
	runSample(t, k, grid, queries, expect)
}

func TestSample2(t *testing.T) {
	k := 8
	grid := [][]int{
		{1, 2, 2, 8},
		{1, 3, 4, 7},
		{5, 1, 7, 6},
		{2, 3, 8, 8},
	}
	queries := [][]int{
		{1, 1, 2, 2},
		{1, 1, 3, 4},
		{1, 1, 2, 4},
		{1, 1, 4, 4},
	}
	expect := []int16{2, 3, 3, 4}
	runSample(t, k, grid, queries, expect)
}
