package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, ranges [][]int, queries []int, expect []int) {
	res := solve(s, ranges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01"
	ranges := [][]int{
		{1, 2},
		{1, 2},
	}
	queries := []int{1, 1, 2, 2}
	expect := []int{0, 1, 0, 1}
	runSample(t, s, ranges, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "10011010"
	ranges := [][]int{
		{5, 6},
		{2, 3},
		{6, 8},
		{5, 7},
		{5, 8},
		{6, 8},
	}
	queries := []int{3, 5, 6, 2, 5, 2, 5, 8, 4, 1}
	expect := []int{2, 3, 2, 2, 1, 2, 2, 2, 2, 2}
	runSample(t, s, ranges, queries, expect)
}
