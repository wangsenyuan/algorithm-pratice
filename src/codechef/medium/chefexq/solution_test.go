package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, A []int, queries [][]int, expect []int) {
	res := solve(n, m, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", n, m, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 3
	A := []int{1, 1, 1, 1, 1}
	queries := [][]int{
		{2, 5, 1},
		{1, 3, 2},
		{2, 5, 1},
	}
	expect := []int{3, 1}
	runSample(t, n, m, A, queries, expect)
}
