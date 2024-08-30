package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int) {
	res := solve(a, queries)

	for i := 0; i < len(queries); i++ {
		v := res[i]
		l, r, x := queries[i][0], queries[i][1], queries[i][2]
		if v < 0 {
			// no answer
			for j := l - 1; j < r; j++ {
				if a[j] != x {
					t.Fatalf("Sample result %v, report no non-%d found between %d %d, but there is one %d", res, x, l, r, j)
				}
			}
		} else {
			if v < l || v > r || a[v-1] == x {
				t.Fatalf("Sample result %v, report a non-%d between %d %d, but it is x", res, x, l, r)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1, 1, 3, 5}
	queries := [][]int{
		{1, 4, 1},
		{2, 6, 2},
		{3, 4, 1},
		{3, 4, 2},
	}
	runSample(t, a, queries)
}

func TestSample2(t *testing.T) {
	a := []int{569888}
	queries := [][]int{
		{1, 1, 967368},
	}
	runSample(t, a, queries)
}
