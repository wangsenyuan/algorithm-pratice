package main

import "testing"

func runSample(t *testing.T, words []int, line_width int, queries [][]int, expect []int) {
	solver := NewSolver(words, line_width)

	var i int
	for _, cur := range queries {
		if len(cur) == 1 {
			id := cur[0]
			ans := solver.QueryLineNo(id)
			if ans != expect[i] {
				t.Fatalf("Sample result for %d, not correct, expect %d, but got %d", id, expect[i], ans)
			}
			i++
		} else {
			id := cur[0]
			w := cur[1]
			solver.ChangeWordLength(id, w)
		}
	}
}

func TestSample1(t *testing.T) {
	m := 4
	words := []int{1, 2}
	queries := [][]int{
		{2},
		{2, 3},
		{2},
	}
	expect := []int{1, 2}
	runSample(t, words, m, queries, expect)
}

func TestSample2(t *testing.T) {
	m := 6
	words := []int{1, 5, 4, 5, 6}
	queries := [][]int{
		{2},
		{4},
		{5},
		{4, 1},
		{4},
		{5},
		{2, 1},
		{3, 2},
		{5, 4},
		{5},
		{3},
	}
	expect := []int{2, 4, 5, 3, 4, 2, 1}
	runSample(t, words, m, queries, expect)
}
