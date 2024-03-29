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
	n := 5
	edges := [][]int{
		{1, 10},
		{1, 1},
		{3, 2},
		{3, 3},
	}
	queries := [][]int{
		{1, 1, 5},
		{5, 4, 5},
		{4, 1, 2},
	}
	expect := []int{3, 0, 13}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 1000000000},
		{2, 1000000000},
		{1, 1000000000},
		{1, 1000000000},
	}
	queries := [][]int{
		{3, 4, 5},
		{2, 1, 5},
		{2, 4, 5},
	}
	expect := []int{3000000000, 1000000000, 2000000000}
	runSample(t, n, edges, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 11
	edges := [][]int{
		{1, 7},
		{2, 1},
		{1, 20},
		{1, 2},
		{5, 6},
		{6, 2},
		{6, 3},
		{5, 1},
		{9, 10},
		{9, 11},
	}
	queries := [][]int{
		{5, 1, 11},
		{1, 1, 4},
		{9, 4, 8},
		{6, 1, 4},
		{9, 7, 11},
		{9, 10, 11},
		{8, 1, 11},
		{11, 4, 5},
	}
	expect := []int{8, 8, 9, 16, 9, 10, 0, 34}
	runSample(t, n, edges, queries, expect)
}
