package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q int, queries [][]int, expect []int) {
	res := solve(n, A, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 1, 2, 2}
	Q := 3
	queries := [][]int{
		{1, 1},
		{2, 3},
		{2, 5},
	}

	expect := []int{1, 0, 3}

	runSample(t, n, A, Q, queries, expect)
}
