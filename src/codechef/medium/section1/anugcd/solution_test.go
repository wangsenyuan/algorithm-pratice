package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, A []int, M int, queries [][]int, expect [][]int) {
	res := solve(N, A, M, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", N, A, M, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 6, 5
	A := []int{1, 2, 3, 4, 5, 4}
	queries := [][]int{
		{2, 1, 5},
		{121, 1, 6},
		{3, 2, 6},
		{5, 5, 5},
		{24, 4, 6},
	}
	expect := [][]int{
		{4, 1},
		{-1, -1},
		{3, 1},
		{5, 1},
		{4, 2},
	}
	runSample(t, N, A, M, queries, expect)
}
