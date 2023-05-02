package p2658

var dd = []int{-1, 0, 1, 0, -1}

func findMaxFish(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(x int, y int) int

	dfs = func(x int, y int) int {
		var res = grid[x][y]
		grid[x][y] = 0
		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] != 0 {
				res += dfs(u, v)
			}
		}

		return res
	}

	var ans int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				tmp := dfs(i, j)
				if tmp > ans {
					ans = tmp
				}
			}
		}
	}
	return ans
}
