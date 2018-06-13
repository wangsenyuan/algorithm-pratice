package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []int64) {
	res := solve(n, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 4, 9, 10}
	queries := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
	}

	expect := []int64{2, 4, 8, 8, 8}
	runSample(t, n, A, queries, expect)
}
