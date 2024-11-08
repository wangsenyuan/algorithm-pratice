package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, k int, queries [][]int, expect []int) {
	res := solve(a, k, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	a := []int{1, 2, 3, 2, 1, 2, 3}
	queries := [][]int{
		{1, 7},
		{2, 7},
		{3, 7},
	}
	expect := []int{6, 5, 2}
	runSample(t, a, k, queries, expect)
}

func TestSample2(t *testing.T) {
	k := 4
	a := []int{4, 3, 1, 1, 2, 4, 3, 2}
	queries := [][]int{
		{3, 6},
		{1, 5},
	}
	expect := []int{2, 5}
	runSample(t, a, k, queries, expect)
}

func TestSample3(t *testing.T) {
	k := 4
	a := []int{4, 5, 1, 2, 3}
	queries := [][]int{
		{1, 4},
		{1, 5},
	}
	expect := []int{2, 3}
	runSample(t, a, k, queries, expect)
}
