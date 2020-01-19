package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int64) {
	var res []int64

	solver := NewSolver(n)

	for _, query := range queries {
		if query[0] == 1 {
			tm, i, sp := query[1], query[2]-1, query[3]
			solver.Change(tm, i, sp)
		} else {
			tm := query[1]
			res = append(res, solver.GetMax(tm))
		}
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	queries := [][]int{
		{1, 1, 1, 2},
		{1, 1, 2, 5},
		{1, 2, 3, 4},
		{1, 2, 4, 7},
		{2, 3},
		{2, 4},
		{1, 5, 5, 4},
		{2, 5},
		{2, 6},
		{1, 7, 5, 8},
		{2, 7},
		{2, 8},
		{2, 9},
		{2, 10},
	}
	expect := []int64{
		10,
		15,
		21,
		28,
		35,
		42,
		49,
		56,
	}
	runSample(t, n, queries, expect)
}
