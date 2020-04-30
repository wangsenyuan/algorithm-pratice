package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, points [][]int, Q int, queries [][]int, expect []int) {
	res := solve(n, points, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, points, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	points := [][]int{
		{1, 2},
		{14, 7},
		{6, 3},
		{8, 7},
		{11, 10},
		{14, 2},
	}
	Q := 3
	queries := [][]int{
		{0, 10},
		{2, 22},
		{11, 17},
	}

	expect := []int{1, 3, 1}

	runSample(t, n, points, Q, queries, expect)
}
