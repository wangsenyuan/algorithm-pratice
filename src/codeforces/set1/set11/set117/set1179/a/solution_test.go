package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries []int, expect [][]int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	queries := []int{1, 2, 10}
	expect := [][]int{
		{1, 2},
		{2, 3},
		{5, 2},
	}
	runSample(t, a, queries, expect)
}

