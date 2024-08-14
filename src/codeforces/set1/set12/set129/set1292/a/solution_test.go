package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []bool) {
	res := solve(n, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	queries := [][]int{
		{2, 3},
		{1, 4},
		{2, 4},
		{2, 3},
		{1, 4},
	}
	expect := []bool{true, false, false, false, true}
	runSample(t, n, queries, expect)
}
