package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect bool) {
	res := solve(len(grid), grid)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n := len(grid)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, 2*n)
		for j := 0; j < 2*n; j++ {
			mat[i][j] = -1
		}
	}

	que := make([]int, n*n)

	bfs := func(x, y int) {
		var head, tail int
		que[head] = x*n + y
		head++
		for tail < head {
			u, v := que[tail]/n, que[tail]%n
			tail++
			for i := 0; i < 4; i++ {
				r, c := u+dd[i], v+dd[i+1]
				if r >= 0 && r < n && c >= 0 && c < n && res[r][c] == D[i] {
					mat[r][2*c] = x + 1
					mat[r][2*c+1] = y + 1
					que[head] = r*n + c
					head++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] == 'X' {
				mat[i][2*j] = i + 1
				mat[i][2*j+1] = j + 1
				bfs(i, j)
			}
		}
	}

	// 在环中的不检查了
	if !reflect.DeepEqual(mat, grid) {
		t.Fatalf("Sample result %v, getting a different grid %v", res, mat)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, 2, 2, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	expect := true
	runSample(t, grid, expect)
}
