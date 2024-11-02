package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v,but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 20, 10}
	queries := [][]int{
		{1, 1},
		{2, 1},
		{2, 2},
		{3, 1},
		{3, 2},
		{3, 3},
	}
	expect := []int{20, 10, 20, 10, 20, 10}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1, 3, 1, 2, 1}
	queries := [][]int{
		{2, 1},
		{2, 2},
		{3, 1},
		{3, 2},
		{3, 3},
		{1, 1},
		{7, 1},
		{7, 7},
		{7, 4},
	}
	expect := []int{2, 3, 2, 3, 2, 3, 1, 1, 3}
	runSample(t, a, queries, expect)
}
