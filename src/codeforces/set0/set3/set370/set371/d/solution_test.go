package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 10}
	queries := [][]int{
		{1, 1, 4},
		{2, 1},
		{1, 2, 5},
		{1, 1, 4},
		{2, 1},
		{2, 2},
	}
	expect := []int{4, 5, 8}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 10, 8}
	queries := [][]int{
		{1, 1, 12},
		{2, 2},
		{1, 1, 6},
		{1, 3, 2},
		{2, 2},
		{2, 3},
	}
	expect := []int{7, 10, 5}
	runSample(t, a, queries, expect)
}
