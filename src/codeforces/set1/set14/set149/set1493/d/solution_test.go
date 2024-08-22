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
	a := []int{1, 6, 8, 12}
	queries := [][]int{
		{1, 12},
		{2, 3},
		{3, 3},
	}
	expect := []int{2, 2, 6}
	runSample(t, a, queries, expect)
}
