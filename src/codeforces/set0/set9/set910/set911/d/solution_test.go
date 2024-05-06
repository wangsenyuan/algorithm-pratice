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
	a := []int{1, 2, 3}
	queries := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := []int{1, 0}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 4, 3}
	queries := [][]int{
		{1, 1},
		{1, 4},
		{1, 4},
		{2, 3},
	}
	expect := []int{1, 1, 1, 0}
	runSample(t, a, queries, expect)
}
