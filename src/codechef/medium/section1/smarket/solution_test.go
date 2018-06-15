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
	n, Q := 8, 5
	A := []int{20, 10, 10, 10, 7, 7, 7, 10}
	queries := [][]int{
		{2, 6, 2},
		{3, 6, 2},
		{3, 6, 3},
		{3, 8, 3},
		{3, 8, 1},
	}

	expect := []int{2, 2, 0, 1, 3}
	runSample(t, n, A, Q, queries, expect)
}

func TestSample2(t *testing.T) {
	n, Q := 3, 2
	A := []int{27, 27, 27}
	queries := [][]int{
		{1, 3, 1},
		{2, 2, 1},
	}

	expect := []int{1, 1}
	runSample(t, n, A, Q, queries, expect)
}
