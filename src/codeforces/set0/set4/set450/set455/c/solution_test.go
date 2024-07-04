package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := solve(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sampple expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	queries := [][]int{
		{2, 1, 2},
		{2, 3, 4},
		{2, 5, 6},
		{2, 3, 2},
		{2, 5, 3},
		{1, 1},
	}
	expect := []int{4}
	runSample(t, n, nil, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{3, 5},
		{1, 6},
		{5, 7},
	}
	queries := [][]int{
		{1, 1},
		{2, 8, 10},
		{1, 10},
		{2, 9, 5},
		{1, 7},
		{1, 4},
		{2, 1, 9},
	}
	expect := []int{5, 1, 5, 5}
	runSample(t, n, edges, queries, expect)
}
