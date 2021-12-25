package main

import "testing"

func runSample(t *testing.T, n, m int, mat [][]int, expect bool) {
	res := solve(n, m, mat)

	if res != expect {
		t.Errorf("sample expect %t, but got %t", expect, res)
	} else if res {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i > 0 && mat[i][j] < mat[i-1][j] {
					t.Errorf("sample result %v is not valid at %d %d", mat, i, j)
					return
				}
				if j > 0 && mat[i][j] < mat[i][j-1] {
					t.Errorf("sample result %v is not valid at %d %d", mat, i, j)
					return
				}
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	mat := [][]int{
		{1, 2, 2, 3},
		{1, -1, 7, -1},
		{6, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	runSample(t, n, m, mat, true)
}
