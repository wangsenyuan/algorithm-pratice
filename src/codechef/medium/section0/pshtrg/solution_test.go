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
	A := []int{3, 1, 8, 9, 7}

	queries := [][]int{
		{2, 1, 5},
		{1, 2, 12},
		{2, 1, 3},
		{2, 2, 5},
	}
	expect := []int64{24, 0, 29}
	runSample(t, n, A, queries, expect)
}
