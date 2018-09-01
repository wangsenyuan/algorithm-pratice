package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, points [][]int, queries [][]int, expect []int) {
	res := solve(N, Q, points, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, points, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 5, 3
	points := [][]int{
		{1, 3},
		{1, 5},
		{3, 6},
		{4, 4},
		{2, 6},
	}
	queries := [][]int{
		{1, 5, 3},
		{1, 5, 4},
		{1, 1, 1},
	}
	expect := []int{3, 3, 0}
	runSample(t, N, Q, points, queries, expect)
}

func TestSample2(t *testing.T) {
	N, Q := 2, 2
	points := [][]int{
		{1, 5},
		{3, 7},
	}
	queries := [][]int{
		{2, 5, 6},
		{2, 3, 4},
	}
	expect := []int{1, 0}
	runSample(t, N, Q, points, queries, expect)
}
