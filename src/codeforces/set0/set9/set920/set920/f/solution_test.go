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
	a := []int{6, 4, 1, 10, 3, 2, 4}
	queries := [][]int{
		{2, 1, 7},
		{2, 4, 5},
		{1, 3, 5},
		{2, 4, 4},
		{1, 5, 7},
		{2, 1, 7},
	}
	expect := []int{30, 13, 4, 22}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 6, 10}
	// 1, 2, 2, 3, 4, 4
	// 1, 2, 2, 3, 3, 4
	// 1, 2, 2, 3, 2, 3
	queries := [][]int{
		{2, 1, 6},
		{1, 3, 6},
		{2, 1, 6},
		{1, 5, 5},
		{2, 1, 6},
		{1, 5, 6},
		{2, 1, 6},
		{1, 1, 6},
		{2, 1, 6},
	}
	expect := []int{26, 16, 15, 13, 11}
	runSample(t, a, queries, expect)
}
