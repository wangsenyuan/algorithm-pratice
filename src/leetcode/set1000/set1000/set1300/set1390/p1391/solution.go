package p1391

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	findValid := func(x, y int) (int, int) {
		if grid[x][y] == 1 {
			u, v := x, y-1
			if v >= 0 && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 4 || grid[u][v] == 6) {
				return u, v
			}
			u, v = x, y+1
			if v < n && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 3 || grid[u][v] == 5) {
				return u, v
			}
		} else if grid[x][y] == 2 {
			u, v := x-1, y
			if u >= 0 && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 4 || grid[u][v] == 3) {
				return u, v
			}
			u, v = x+1, y
			if u < m && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 5 || grid[u][v] == 6) {
				return u, v
			}
		} else if grid[x][y] == 3 {
			u, v := x, y-1
			if v >= 0 && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 4 || grid[u][v] == 6) {
				return u, v
			}
			u, v = x+1, y
			if u < m && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 5 || grid[u][v] == 6) {
				return u, v
			}
		} else if grid[x][y] == 4 {
			u, v := x, y+1
			if v < n && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 3 || grid[u][v] == 5) {
				return u, v
			}
			u, v = x+1, y
			if u < m && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 5 || grid[u][v] == 6) {
				return u, v
			}
		} else if grid[x][y] == 5 {
			u, v := x-1, y
			if u >= 0 && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 4 || grid[u][v] == 3) {
				return u, v
			}
			u, v = x, y-1
			if v >= 0 && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 4 || grid[u][v] == 6) {
				return u, v
			}
		} else {
			u, v := x-1, y
			if u >= 0 && !vis[u][v] && (grid[u][v] == 2 || grid[u][v] == 4 || grid[u][v] == 3) {
				return u, v
			}
			u, v = x, y+1
			if v < n && !vis[u][v] && (grid[u][v] == 1 || grid[u][v] == 3 || grid[u][v] == 5) {
				return u, v
			}
		}

		return -1, -1
	}

	var x, y int

	if grid[0][0] == 1 || grid[0][0] == 2 || grid[0][0] == 3 || grid[0][0] == 6 {
		vis[0][0] = true
		for x != m-1 || y != n-1 {
			x, y = findValid(x, y)
			if x < 0 || x >= m || y < 0 || y >= n {
				return false
			}
			vis[x][y] = true
		}
		return x == m-1 && y == n-1
	} else if grid[0][0] == 4 {
		vis[0][0] = true
		if y < n-1 {
			vis[0][1] = true
			x, y = 0, 1
			for x != m-1 || y != n-1 {
				x, y = findValid(x, y)
				if x < 0 || x >= m || y < 0 || y >= n {
					break
				}
				vis[x][y] = true
			}
			if x == m-1 && y == n-1 {
				return true
			}
		}

		if x < m-1 {
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					vis[i][j] = false
				}
			}
			vis[0][0] = true
			vis[1][0] = true
			x, y = 1, 0
			for x != m-1 || y != n-1 {
				x, y = findValid(x, y)
				if x < 0 || x >= m || y < 0 || y >= n {
					return false
				}
				vis[x][y] = true
			}
			if x == m-1 && y == n-1 {
				return true
			}
		}

	}

	return x == m-1 && y == n-1
}
