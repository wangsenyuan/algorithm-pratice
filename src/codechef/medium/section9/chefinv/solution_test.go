package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, queries [][]int, expect []int64) {
	res := solve(A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 3, 3, 2, 5}
	queries := [][]int{
		{1, 1},
		{1, 3},
		{2, 5},
	}
	expect := []int64{5, 6, 0}
	runSample(t, A, queries, expect)
}
