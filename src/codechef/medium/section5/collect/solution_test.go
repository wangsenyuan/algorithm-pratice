package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, A []int, Q int, queries [][]int, expect []int) {
	res := solve(N, A, Q, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", N, A, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 8
	A := []int{2, 1, 1, 2, 1, 2, 2, 2}
	Q := 5
	queries := [][]int{
		{1, 2, 2},
		{2, 1, 8},
		{2, 5, 7},
		{2, 3, 3},
		{2, 4, 4},
	}
	expect := []int{2065880, 90, 1, 1}
	runSample(t, N, A, Q, queries, expect)
}
