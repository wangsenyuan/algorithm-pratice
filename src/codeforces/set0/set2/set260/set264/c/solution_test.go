package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, v []int, c []int, queries [][]int, expect []int) {
	res := solve(v, c, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	v := []int{1, -2, 3, 4, 0, -1}
	c := []int{1, 2, 1, 2, 1, 1}
	queries := [][]int{
		{5, 1},
		{-2, 1},
		{1, 0},
	}
	expect := []int{20, 9, 4}
	runSample(t, v, c, queries, expect)
}

func TestSample2(t *testing.T) {
	v := []int{-3, 6, -1, 2}
	c := []int{1, 2, 3, 1}
	queries := [][]int{
		{1, -1},
	}
	expect := []int{5}
	runSample(t, v, c, queries, expect)
}
