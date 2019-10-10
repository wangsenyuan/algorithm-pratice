package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, ans []bool) {
	res := solve(n, queries)

	if !reflect.DeepEqual(res, ans) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, queries, ans, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	queries := [][]int{
		{2, 0, 1},
		{1, 3, 1},
		{2, 2, 4},
		{2, 5, 6},
		{2, 8, 8},
		{2, 9, 10},
	}

	expect := []bool{false, true, true, false, true, false}

	runSample(t, n, queries, expect)
}
