package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []int) {
	solve(n, A, queries)

	if !reflect.DeepEqual(A, expect) {
		t.Errorf("sample expect %v, but got %v", expect, A)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	queries := [][]int{
		{1, 2, 2, 2},
		{1, 2, 2, 2},
		{2, 2, 3},
		{1, 2, 3, 3},
		{2, 1, 5},
	}

	expect := []int{5, 1, 1}
	runSample(t, n, A, queries, expect)
}
