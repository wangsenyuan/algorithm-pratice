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
	a := []int{1, 3, 0, 2, 1}
	queries := [][]int{
		{1, 3},
		{2, 4},
		{1, 4},
	}
	expect := []bool{false, true, true}
	runSample(t, a, queries, expect)
}
