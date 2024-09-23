package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A int, B int, queries [][]int, expect []int) {
	res := solve(A, B, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B := 2, 1
	queries := [][]int{
		{1, 5, 3},
		{3, 3, 10},
		{7, 10, 2},
		{6, 4, 8},
	}
	expect := []int{4, -1, 8, -1}

	runSample(t, A, B, queries, expect)
}

func TestSample2(t *testing.T) {
	A, B := 1, 5
	queries := [][]int{
		{1, 5, 10},
		{2, 7, 4},
	}
	expect := []int{1, 2}

	runSample(t, A, B, queries, expect)
}
