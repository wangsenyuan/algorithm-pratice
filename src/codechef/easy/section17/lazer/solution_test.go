package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []int) {
	res := solve(n, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 3, 5, 1}
	queries := [][]int{
		{2, 4, 4},
		{1, 2, 3},
		{1, 4, 1},
	}
	expect := []int{2, 1, 2}
	runSample(t, n, A, queries, expect)
}
