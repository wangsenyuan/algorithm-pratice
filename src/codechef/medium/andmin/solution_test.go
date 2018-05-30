package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, q int, queries [][]int, expect []int) {
	res := solve(n, A, q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 5, 5
	A := []int{1, 5, 2, 3, 4}
	queries := [][]int{
		{0, 2, 5},
		{1, 1, 5, 6},
		{0, 2, 2},
		{1, 2, 5, 3},
		{0, 1, 3},
	}
	expect := []int{2, 4, 0}
	runSample(t, n, A, q, queries, expect)
}
