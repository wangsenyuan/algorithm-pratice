package main

import "testing"

func runSample(t *testing.T, n int, s int, nums []int, queries [][]int, expect []int) {
	solver := NewSolver1(n, s, nums)

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		x, y := cur[0], cur[1]
		res := solver.Query(x, y)
		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d at %d", expect[i], res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	s := 0
	nums := []int{4, 2, 3, 1, 5, 5}
	queries := [][]int{
		{1, 6},
		{2, 4},
		{3, 4},
	}
	expect := []int{3, 2, 1}
	runSample(t, n, s, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	s := 1
	nums := []int{7, 8, 8, 9, 10, 10, 2, 3, 5, 10}
	queries := [][]int{
		{9, 6},
		{2, 9},
		{10, 6},
		{4, 2},
		{5, 10},
		{4, 5},
		{8, 7},
		{8, 2},
		{7, 7},
		{9, 5},
	}
	expect := []int{3, 3, 3, 1, 4, 2, 2, 4, 1, 4}
	runSample(t, n, s, nums, queries, expect)
}
