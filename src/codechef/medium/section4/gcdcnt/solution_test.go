package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q int, queries [][]int, expect []int) {
	res := solve(n, A, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{2, 3, 4, 5}
	Q := 3
	queries := [][]int{
		{2, 1, 4, 2},
		{1, 2, 6},
		{2, 1, 4, 2},
	}
	expect := []int{2, 1}
	runSample(t, n, A, Q, queries, expect)
}
