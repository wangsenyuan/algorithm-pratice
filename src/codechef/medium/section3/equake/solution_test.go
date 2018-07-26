package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, m int, queries [][]int, expect []int) {
	res := solve(n, A, m, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, m, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{17, 3140, 832}
	m := 8
	queries := [][]int{
		{1, 0, 2},
		{0, 0, 2, 1},
		{1, 1, 2},
		{1, 0, 0},
		{0, 0, 2, 2},
		{1, 0, 2},
		{0, 1, 1, 1},
		{1, 0, 2},
	}
	expect := []int{3140, 1403, 71, 832, 3140}
	runSample(t, n, A, m, queries, expect)
}
