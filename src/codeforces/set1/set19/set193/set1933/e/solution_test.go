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
	a := []int{3, 1, 4, 1, 5, 9}
	queries := [][]int{
		{1, 8},
		{2, 7},
		{5, 9},
	}
	expect := []int{3, 4, 5}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10}
	queries := [][]int{
		{1, 1},
	}
	expect := []int{1}
	runSample(t, a, queries, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 10, 9, 6, 8, 3, 10, 7, 3}
	queries := [][]int{
		{8, 56},
		{1, 12},
		{9, 3},
		{1, 27},
		{5, 45},
	}
	expect := []int{9, 2, 9, 4, 9}
	runSample(t, a, queries, expect)
}

func TestSample4(t *testing.T) {
	a := []int{7, 9, 2, 5, 2}
	queries := [][]int{
		{1, 37},
		{2, 9},
		{3, 33},
		{4, 32},
		{4, 15},
		{2, 2},
		{4, 2},
		{2, 19},
		{3, 7},
		{2, 7},
	}
	expect := []int{5, 2, 5, 5, 5, 2, 4, 5, 4, 2}
	runSample(t, a, queries, expect)
}
