package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries [][]int, expect []int) {
	res := solve(queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := [][]int{
		{1, 3, 4},
		{2, 3, 0},
		{2, 4, 3},
		{1, 4, -4},
		{2, 1, 0},
	}
	expect := []int{4, 4, 0}
	runSample(t, queries, expect)
}
