package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a int, b int, queries [][]int, expect []int) {
	res := solve(a, b, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 4, 6
	queries := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{1, 9},
	}
	expect := []int{0, 0, 0, 2, 4}
	runSample(t, a, b, queries, expect)
}
