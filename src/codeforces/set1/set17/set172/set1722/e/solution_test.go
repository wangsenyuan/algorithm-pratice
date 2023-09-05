package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rects [][]int, queries [][]int, expect []int64) {
	res := solve(rects, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rects := [][]int{
		{2, 3},
		{3, 2},
	}
	queries := [][]int{
		{1, 1, 3, 4},
	}
	expect := []int64{6}
	runSample(t, rects, queries, expect)
}

func TestSample2(t *testing.T) {
	rects := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}
	queries := [][]int{
		{3, 3, 6, 6},
		{2, 1, 4, 5},
		{1, 1, 2, 10},
		{1, 1, 100, 100},
		{1, 1, 3, 3},
	}
	expect := []int64{41, 9, 0, 54, 4}
	runSample(t, rects, queries, expect)
}

func TestSample3(t *testing.T) {
	rects := [][]int{
		{999, 999},
		{999, 999},
		{999, 998},
	}
	queries := [][]int{
		{1, 1, 1000, 1000},
	}
	expect := []int64{2993004}
	runSample(t, rects, queries, expect)
}
