package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	var visit func(i, j int) int
	dir := [...]int{-1, 0, 1, 0, -1}

	visit = func(i, j int) int {
		if grid[i][j] != 1 {
			return 0
		}

		grid[i][j] = -1

		tmp := 1

		for k := 0; k < 4; k++ {
			x, y := i+dir[k], j+dir[k+1]
			if x >= 0 && x < n && y >= 0 && y < m {
				tmp += visit(x, y)
			}
		}

		return tmp
	}
	var ans = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				tmp := visit(i, j)
				if tmp > ans {
					ans = tmp
				}
			}
		}
	}

	return ans
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
