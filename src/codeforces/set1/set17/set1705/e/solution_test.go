package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, queries [][]int, expect []int) {
	res := solve(A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2, 2, 4, 5}
	queries := [][]int{
		{2, 3},
		{5, 3},
		{4, 1},
		{1, 4},
	}
	expect := []int{6, 5, 4, 5}
	runSample(t, A, queries, expect)
}

func TestSample2(t *testing.T) {
	A := []int{200000, 1}
	queries := [][]int{
		{2, 200000},
	}
	expect := []int{200001}
	runSample(t, A, queries, expect)
}
