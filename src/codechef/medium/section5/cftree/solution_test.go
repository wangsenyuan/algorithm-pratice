package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M int, parent []int, queries [][]int, expect []int) {
	res := solve(N, M, parent, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, M, parent, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 10
	parent := []int{1, 1, 3, 3}
	queries := [][]int{
		{1, 3},
		{0, 1, 2},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{0, 3, 1},
		{1, 3},
		{1, 4},
		{1, 5},
	}
	expect := []int{0, 2, 2, 3, 3, 3, 4, 4}

	runSample(t, N, M, parent, queries, expect)
}
