package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, A []int, queries [][]int, expect [][]int) {
	res := solve(n, m, A, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", n, m, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 10
	A := []int{6, 8, 10, 13}
	queries := [][]int{
		{3, 1, 2, 2},
		{2, 1, 2},
		{3, 1, 2, 3},
		{2, 2, 3},
		{1, 1, 7},
		{3, 1, 3, 10},
		{2, 3, 1},
		{3, 1, 3, 2},
		{2, 1, 4},
		{3, 1, 4, 6},
	}
	expect := [][]int{
		{0, 0},
		{-9, 4},
		{7, 1},
		{0, 0},
		{0, 0},
	}
	runSample(t, n, m, A, queries, expect)
}
