package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, E [][]int, queries [][]int, expect []int64) {
	res := solve(n, A, E, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %v, expect %v, but got %v", n, A, E, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	A := []int{2, 6, 4, 3, 5}
	queries := [][]int{
		{1, 4},
		{2, 2},
	}
	expect := []int64{9, 4}

	runSample(t, n, A, E, queries, expect)
}
