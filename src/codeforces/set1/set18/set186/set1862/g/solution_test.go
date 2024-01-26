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
	a := []int{2, 4, 8}
	queries := [][]int{
		{1, 6},
		{2, 10},
		{3, 1},
	}
	expect := []int{10, 12, 15}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2, 2, 2}
	queries := [][]int{
		{5, 3},
	}
	expect := []int{4}
	runSample(t, a, queries, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 6}
	queries := [][]int{
		{1, 2},
		{1, 7},
		{1, 7},
		{2, 5},
		{1, 2},
		{2, 7},
		{2, 2},
	}
	expect := []int{10, 8, 8, 9, 8, 12, 2}
	runSample(t, a, queries, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 5, 1, 10, 6}
	queries := [][]int{
		{1, 7},
		{4, 8},
		{2, 5},
		{1, 4},
		{2, 8},
		{3, 4},
		{1, 9},
		{3, 7},
		{3, 4},
		{3, 1},
	}
	expect := []int{14, 12, 12, 11, 11, 10, 11, 10, 11, 14}
	runSample(t, a, queries, expect)
}
