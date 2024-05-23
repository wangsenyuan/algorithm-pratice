package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []bool) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1, 4, 5}
	queries := [][]int{
		{1, 5},
		{4, 4},
		{3, 4},
		{1, 3},
	}
	expect := []bool{true, false, true, false}
	runSample(t, a, queries, expect)
}
