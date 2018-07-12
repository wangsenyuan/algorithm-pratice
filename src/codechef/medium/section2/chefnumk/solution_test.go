package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, A []int, queries [][]int, expect []int) {
	res := solve(N, Q, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 5, 3
	A := []int{1, 2, 3, 2, 1}
	queries := [][]int{
		{2, 4, 2},
		{1, 5, 1},
		{1, 5, 2},
	}
	expect := []int{2, 1, 3}
	runSample(t, N, Q, A, queries, expect)
}
