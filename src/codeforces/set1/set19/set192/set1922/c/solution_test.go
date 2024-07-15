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
	a := []int{0, 8, 12, 15, 20}
	queries := [][]int{
		{1, 4},
		{1, 5},
		{3, 4},
		{3, 2},
		{5, 1},
	}
	expect := []int{3, 8, 1, 4, 14}
	runSample(t, a, queries, expect)
}
