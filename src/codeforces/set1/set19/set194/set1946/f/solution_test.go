package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 6, 3, 5, 4, 8, 7}
	queries := [][]int{
		{1, 8},
		{2, 8},
		{1, 7},
		{1, 6},
		{1, 3},
		{5, 8},
		{4, 4},
		{2, 3},
	}
	expect := []int{20, 15, 18, 12, 5, 5, 1, 3}
	runSample(t, a, queries, expect)
}
