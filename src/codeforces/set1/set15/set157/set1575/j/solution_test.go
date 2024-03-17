package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat [][]int, queries []int, expect []int) {
	res := solve(mat, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 2, 3, 3, 3},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
	}
	queries := []int{1, 2, 1}
	expect := []int{2, 2, 1}
	runSample(t, mat, queries, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 3},
	}
	queries := []int{1, 2}
	expect := []int{1, 2}
	runSample(t, mat, queries, expect)
}
