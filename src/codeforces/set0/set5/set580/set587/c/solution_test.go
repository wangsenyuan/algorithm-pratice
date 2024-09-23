package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, c []int, queries [][]int, expect [][]int) {
	res := solve(n, edges, c, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 3},
		{1, 2},
		{1, 4},
		{4, 5},
	}
	c := []int{2, 1, 4, 3}
	queries := [][]int{
		{4, 5, 6},
		{1, 5, 2},
		{5, 5, 10},
		{2, 3, 3},
		{5, 3, 1},
	}
	expect := [][]int{
		{3},
		{2, 3},
		{},
		{1, 2, 4},
		{2},
	}
	runSample(t, n, edges, c, queries, expect)
}
