package p1559

var dd = []int{-1, 0, 1, 0, -1}

func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])
	if m < 2 || n < 2 {
		return false
	}
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	var dfs func(x, y, px, py int) bool

	dfs = func(x, y, px, py int) bool {
		if vis[x][y] == 1 {
			return true
		}
		if vis[x][y] == 2 {
			// this path is already processed
			return false
		}
		vis[x][y] = 1
		// dir = TOP(0), RIGHT(1), BOTTOM(2), LEFT(3)
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n &&
				(px != u || py != v) &&
				grid[u][v] == grid[x][y] && dfs(u, v, x, y) {
				return true
			}
		}

		vis[x][y] = 2
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] == 0 && dfs(i, j, -1, -1) {
				return true
			}
		}
	}

	return false
}
