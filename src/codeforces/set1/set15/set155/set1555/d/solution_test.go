package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "baacb"
	queries := [][]int{
		{1, 3},
		{1, 5},
		{4, 5},
		{2, 3},
	}
	expect := []int{1, 2, 0, 1}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "cbcccaacaa"
	queries := [][]int{
		{2, 7},
		{2, 4},
		{5, 5},
		{3, 6},
		{2, 5},
		{5, 10},
		{4, 8},
		{6, 6},
		{1, 4},
		{2, 3},
		{1, 9},
		{4, 6},
		{3, 10},
	}
	expect := []int{3, 1, 0, 2, 2, 2, 2, 0, 1, 0, 4, 1, 4}
	runSample(t, s, queries, expect)
}
