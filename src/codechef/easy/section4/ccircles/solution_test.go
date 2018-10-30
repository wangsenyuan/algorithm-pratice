package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, points [][]int, queries []int, expect []int) {
	res := solve(N, Q, points, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, points, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 2, 3
	points := [][]int{
		{0, 0, 5},
		{8, 3, 2},
	}
	queries := []int{0, 10, 20}
	expect := []int{0, 1, 0}
	runSample(t, N, Q, points, queries, expect)
}
