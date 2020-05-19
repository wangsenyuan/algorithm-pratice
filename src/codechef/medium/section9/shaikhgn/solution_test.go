package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, G [][]int, queries [][]int, expect [][]int) {
	solver := NewSolver(N, G)

	for i := 0; i < len(queries); i++ {
		k, x := queries[i][0], queries[i][1]
		res := solver.Solve(k, x)

		if !reflect.DeepEqual(res, expect[i]) {
			t.Fatalf("Sample expect %v, but got %v", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	N := 5
	G := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	queries := [][]int{
		{3, 1},
		{10000, 1},
		{0, 5},
		{1, 5},
	}
	expect := [][]int{
		{1, 4},
		{2, 4},
		{5},
		{},
	}
	runSample(t, N, G, queries, expect)
}
