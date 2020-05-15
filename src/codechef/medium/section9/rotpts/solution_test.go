package main

import "testing"

func runSample(t *testing.T, n int, P [][]int, A []int, queries [][]int, expect [][]int) {
	solver := NewSolver(n, P, A)

	var i int

	for j := 0; j < len(queries); j++ {
		query := queries[j]
		if query[0] == 1 {
			// query
			x, y, l, r := query[1], query[2], query[3], query[4]
			xx, yy := solver.FindAnswer(x, y, l, r)
			if expect[i][0] != xx || expect[i][1] != yy {
				t.Fatalf("Sample Fail at step %d, expect %v, but got %d %d", i, expect[i], xx, yy)
			}
			i++
			continue
		}
		//update
		u, x, y, a := query[1], query[2], query[3], query[4]
		solver.UpdatePos(u, x, y, a)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := [][]int{{0, 0}, {1, 2}, {3, 2}}
	A := []int{90, 180, 270}
	queries := [][]int{
		{1, 5, 5, 1, 3},
		{2, 2, 2, 2, 90},
		{1, 5, 5, 1, 3},
	}
	expect := [][]int{
		{0, 1000000005},
		{1000000003, 6},
	}
	runSample(t, n, P, A, queries, expect)
}
