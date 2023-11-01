package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat [][]int, queries [][]int, expect []bool) {
	res := solve(mat, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 2, 3, 5},
		{3, 1, 3, 2},
		{4, 5, 2, 3},
		{5, 5, 3, 2},
		{4, 4, 3, 4},
	}
	queries := [][]int{
		{1, 1},
		{2, 5},
		{4, 5},
		{3, 5},
		{1, 3},
		{1, 5},
	}
	expect := []bool{true, false, true, true, true, false}
	runSample(t, mat, queries, expect)
}
