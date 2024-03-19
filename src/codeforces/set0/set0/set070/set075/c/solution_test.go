package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a int, b int, queries [][]int, expect []int) {
	res := solve(a, b, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 9, 27
	queries := [][]int{
		{1, 5},
		{10, 11},
		{9, 11},
	}
	expect := []int{3, -1, 9}
	runSample(t, a, b, queries, expect)
}

func TestSample2(t *testing.T) {
	a, b := 90, 36
	queries := [][]int{
		{13, 15},
	}
	expect := []int{-1}
	runSample(t, a, b, queries, expect)
}
