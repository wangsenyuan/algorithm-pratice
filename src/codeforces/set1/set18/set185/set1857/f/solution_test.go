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
	a := []int{1, 3, 2}
	queries := [][]int{
		{3, 2},
		{5, 6},
		{3, 1},
		{5, 5},
	}
	expect := []int{1, 1, 0, 0}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1}
	queries := [][]int{
		{2, 1},
	}
	expect := []int{6}
	runSample(t, a, queries, expect)
}
