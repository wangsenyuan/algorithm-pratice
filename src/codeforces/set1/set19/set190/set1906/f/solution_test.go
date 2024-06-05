package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, operations [][]int, queries [][]int, expect []int) {
	res := solve(n, operations, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	operations := [][]int{
		{1, 1, -50},
		{1, 2, -20},
		{2, 2, -30},
		{1, 1, 60},
		{1, 2, 40},
		{2, 2, 10},
	}
	queries := [][]int{
		{1, 1, 6},
		{2, 1, 6},
		{1, 1, 3},
		{2, 1, 3},
		{1, 1, 2},
	}
	expect := []int{100, 50, 0, 0, -20}

	runSample(t, n, operations, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	operations := [][]int{
		{1, 3, 3},
		{2, 4, -2},
		{3, 5, 3},
	}
	queries := [][]int{
		{1, 1, 3},
		{2, 1, 3},
		{3, 1, 3},
		{3, 2, 3},
		{2, 2, 3},
		{2, 2, 2},
	}
	expect := []int{3, 3, 4, 3, 0, -2}

	runSample(t, n, operations, queries, expect)
}
