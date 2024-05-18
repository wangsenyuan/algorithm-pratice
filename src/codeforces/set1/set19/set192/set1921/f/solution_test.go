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
	a := []int{1, 1, 2}
	queries := [][]int{
		{1, 2, 2},
		{2, 2, 1},
		{1, 1, 2},
	}
	expect := []int{5, 1, 3}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{34, 87, 5, 42, -44, 66, -32}
	queries := [][]int{
		{2, 2, 2},
		{4, 3, 1},
		{1, 3, 2},
		{6, 2, 1},
		{5, 2, 2},
		{2, 5, 2},
		{6, 1, 2},
	}
	expect := []int{171, 42, 118, 66, -108, 23, 2}
	runSample(t, a, queries, expect)
}
