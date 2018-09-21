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
	n, q := 10, 3
	A := []int{5, 5, 1, 1, 5, 5, 1, 1, 5, 5}
	queries := [][]int{
		{1, 10},
		{2, 9},
		{1, 1},
	}

	expect := []int{8, 8, 0}

	runSample(t, n, A, q, queries, expect)
}
