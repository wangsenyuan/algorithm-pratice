package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect [][]int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 2, 8, 2}
	queries := [][]int{
		{1, 1},
		{0, 1, 3},
		{1, 1},
		{0, 3, 4},
		{1, 2},
	}
	expect := [][]int{
		{8, 7},
		{8, 5},
		{7, 3},
	}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 1, 2, 4, 1, 7, 3, 8, 10, 8}
	queries := [][]int{
		{0, 5, 6},
		{1, 8},
		{1, 1},
		{0, 10, 3},
		{1, 5},
		{1, 3},
		{1, 2},
		{0, 6, 1},
		{1, 9},
		{1, 1},
	}
	expect := [][]int{
		{8, 1},
		{6, 2},
		{5, 1},
		{5, 2},
		{5, 3},
		{9, 1},
		{10, 4},
	}
	runSample(t, a, queries, expect)
}
