package p1568

var dd = []int{-1, 0, 1, 0, -1}

func minDays(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}

	var tot int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tot += grid[i][j]
		}
	}

	if tot <= 1 {
		return 0
	}

	var dfs func(x, y int, color int) int

	dfs = func(x, y int, color int) int {
		var cnt = 1
		vis[x][y] = color
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && vis[u][v] < color && grid[u][v] == 1 {
				cnt += dfs(u, v, color)
			}
		}
		return cnt
	}
	var color = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tmp := dfs(i, j, color)
				color++
				if tmp != tot {
					return 0
				}
			}
		}
	}

	// only one land, best answer could be 2, check where we could find answer 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// mark it as removed
				grid[i][j] = 0

				for k := 0; k < 4; k++ {
					x, y := i+dd[k], j+dd[k+1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
						tmp := dfs(x, y, color)
						color++
						if tmp+1 < tot {
							return 1
						}
					}
				}

				grid[i][j] = 1
			}
		}
	}

	return 2
}
