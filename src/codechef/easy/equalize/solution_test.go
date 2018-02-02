package main

import "testing"

func TestSample(t *testing.T) {
	n, m, q := 3, 7, 3

	rect := [][]int{
		{8, 5, 5, 5, 8, 6, 8},
		{8, 9, 5, 5, 5, 9, 8},
		{8, 6, 8, 5, 5, 5, 8},
	}
	qs := [][]int{
		{3, 3, 8},
		{3, 5, 5},
		{3, 7, 6},
	}
	cols := make([][]int, n)
	for i := 0; i < n; i++ {
		cols[i] = make([]int, m)
	}

	for i := 0; i < q; i++ {
		k, l, expect := qs[i][0], qs[i][1], qs[i][2]
		res := solve(n, m, rect, cols, k, l)
		if res != expect {
			t.Errorf("sample %v %d %d, expect %d, but got %d", rect, k, l, expect, res)
		}
	}
}
