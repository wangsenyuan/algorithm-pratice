package main

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, n int, pairs [][]int, queries [][]int, expect []int) {
	res := solve(n, pairs, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v, expect %v, but got %v", n, pairs, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	pairs := [][]int{
		{1, 2}, {2, 3}, {1, 4},
	}
	queries := [][]int{
		{1, 4, 2},
		{2, 4, 2},
	}

	expect := []int{1, 2}

	runSample(t, n, pairs, queries, expect)
}
