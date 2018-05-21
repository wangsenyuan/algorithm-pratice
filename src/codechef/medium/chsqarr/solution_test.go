package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M int, MTR [][]int, Q int, queries [][]int, expect []int) {
	res := solve(N, M, MTR, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %d %v, expect %v, but got %v", N, M, MTR, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 3, 4
	MTR := [][]int{
		{1, 8, 3, 4},
		{5, 2, 3, 1},
		{3, 6, 2, 2},
	}
	Q := 4
	queries := [][]int{
		{1, 1},
		{2, 2},
		{2, 3},
		{3, 2},
	}
	expect := []int{0, 4, 15, 9}
	runSample(t, N, M, MTR, Q, queries, expect)
}
