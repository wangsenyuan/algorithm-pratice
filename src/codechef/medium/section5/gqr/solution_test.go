package main

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	n := 4
	A := []int{4, 2, 1, 8}
	solver := NewSolver(n, A)

	queries := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := []int{2,
		2,
		4,
		1,
		2,
		1}

	for i := 0; i < len(queries); i++ {
		res := solver.Query(queries[i][0], queries[i][1])
		if res != expect[i] {
			t.Errorf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{4, 2, 1, 8}

	queries := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := []int{2,
		2,
		4,
		1,
		2,
		1}

	res := solve(n, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}
