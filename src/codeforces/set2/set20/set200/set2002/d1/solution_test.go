package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, a []int, p []int, queries [][]int, expect []bool) {
	res := solve(n, a, p, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1, 1}
	p := []int{1, 2, 3}
	queries := [][]int{
		{2, 3},
		{3, 2},
		{1, 3},
	}
	expect := []bool{true, true, false}

	runSample(t, n, a, p, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	a := []int{1, 1, 2, 2, 3, 3}
	p := []int{1, 2, 3, 4, 5, 6, 7}
	queries := [][]int{
		{3, 5},
		{2, 5},
		{3, 7},
		{4, 6},
	}
	expect := []bool{true, false, false, true}

	runSample(t, n, a, p, queries, expect)
}
