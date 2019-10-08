package p1219

var dd = []int{-1, 0, 1, 0, -1}

func getMaximumGold(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(x int, y int) int

	dfs = func(x int, y int) int {
		var ans int

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] > 0 {
				bak := grid[u][v]
				grid[u][v] = 0
				tmp := dfs(u, v) + bak
				grid[u][v] = bak
				ans = max(ans, tmp)
			}
		}

		return ans
	}

	var best int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				bak := grid[i][j]
				grid[i][j] = 0
				ans := dfs(i, j) + bak
				grid[i][j] = bak
				best = max(ans, best)
			}

		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
