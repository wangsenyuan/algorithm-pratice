package main

import "testing"

func TestSample1(t *testing.T) {
	n := 10
	A := []int{9, 17, 93, 16, 3, 61, 23, 11, 2, 1}
	E := [][]int{
		{1, 2},
		{2, 5},
		{5, 8},
		{1, 3},
		{1, 4},
		{3, 6},
		{3, 7},
		{6, 9},
		{6, 10},
	}

	queries := [][]int{
		{4, 14},
		{7, 123},
		{5, 103},
		{9, 32},
		{5, 118},
	}

	expect := [][]int{
		{4, 30},
		{7, 114},
		{8, 30},
		{3, 99},
		{6, 40},
	}

	solver := NewSolver(n, A, E)

	var v, k int

	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		v ^= a
		k ^= b
		v, k = solver.Query(v, k)
		c, d := expect[i][0], expect[i][1]

		if v != c || d != k {
			t.Fatalf("sample fail at %d, expect %v, but got %d %d", i, expect[i], v, k)
		}
	}

}
