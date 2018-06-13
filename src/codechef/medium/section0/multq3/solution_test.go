package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int) {
	res := solve(n, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	queries := [][]int{
		{1, 0, 3},
		{0, 1, 2},
		{0, 1, 3},
		{1, 0, 0},
		{0, 0, 3},
		{1, 3, 3},
		{1, 0, 3},
	}
	expect := []int{4, 1, 0, 2}
	runSample(t, n, queries, expect)
}
