package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, H []int, roads [][]int, queries [][]int, expect []bool) {
	res := solve(H, roads, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{1, 5, 3, 4, 2, 4, 1}
	roads := [][]int{
		{1, 4},
		{4, 3},
		{3, 6},
		{3, 2},
		{2, 5},
		{5, 6},
		{5, 7},
	}
	queries := [][]int{
		{1, 1, 3},
		{6, 2, 0},
		{4, 7, 0},
		{1, 7, 4},
		{1, 7, 2},
	}

	expect := []bool{true, false, true, true, false}

	runSample(t, H, roads, queries, expect)
}

func TestSample2(t *testing.T) {
	H := []int{4, 7, 6, 2, 5, 1}
	roads := [][]int{
		{1, 3},
		{5, 3},
		{1, 5},
		{2, 4},
		{6, 2},
	}
	queries := [][]int{
		{1, 5, 1},
		{1, 3, 1},
		{1, 2, 1000},
		{6, 2, 6},
		{6, 2, 5},
	}

	expect := []bool{true, false, false, true, false}

	runSample(t, H, roads, queries, expect)
}

func TestSample3(t *testing.T) {
	H := []int{7, 9, 2, 10, 8, 6}
	roads := [][]int{
		{4, 2},
		{6, 1},
		{4, 5},
		{3, 5},
		{6, 4},
		{1, 3},
		{2, 6},
		{6, 5},
		{1, 2},
		{3, 6},
	}
	queries := [][]int{
		{4, 4, 8},
		{3, 3, 1},
		{5, 5, 9},
		{2, 1, 7},
		{6, 6, 10},
	}

	expect := []bool{true, true, true, true, true}

	runSample(t, H, roads, queries, expect)
}
