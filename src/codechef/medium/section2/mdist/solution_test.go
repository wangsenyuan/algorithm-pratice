package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, points [][]int, q int, queries [][]int, expect []int64) {
	res := solve(n, points, q, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, points, q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	points := [][]int{
		{10, 10},
		{10, 20},
		{20, 20},
		{20, 10},
	}
	q := 2
	queries := [][]int{
		{0, 1, 20, 20},
		{1, 0, 3},
	}
	expect := []int64{20}
	runSample(t, n, points, q, queries, expect)
}
