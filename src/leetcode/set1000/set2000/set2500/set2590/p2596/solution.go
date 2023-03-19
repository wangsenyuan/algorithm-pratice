package p2596

import "sort"

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	type Move struct {
		r int
		c int
		i int
	}

	n := len(grid)
	move := make([]Move, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			move[i*n+j] = Move{i, j, grid[i][j]}
		}
	}

	sort.Slice(move, func(i, j int) bool {
		return move[i].i < move[j].i
	})

	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true

	x, y := 0, 0

	for i := 1; i < len(move); i++ {
		u, v := move[i].r, move[i].c
		if vis[u][v] {
			return false
		}
		vis[u][v] = true
		if abs(u-x) == 2 && abs(v-y) == 1 {
			x, y = u, v
			continue
		}

		if abs(u-x) == 1 && abs(v-y) == 2 {
			x, y = u, v
			continue
		}
		return false
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
