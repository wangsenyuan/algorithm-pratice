package p980

var dd = [...]int{-1, 0, 1, 0, -1}

func uniquePathsIII(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}

	var x, y, p, q int
	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x, y = i, j
			}

			if grid[i][j] == 2 {
				p, q = i, j
			}

			if grid[i][j] != -1 {
				cnt++
			}
		}
	}

	var res int
	var dfs func(u, v int, sum int)
	dfs = func(u, v int, sum int) {
		if vis[u][v] == 1 {
			return
		}
		if u == p && v == q {
			//reach dest
			if sum == cnt {
				res++
			}
			return
		}

		vis[u][v] = 1

		for k := 0; k < 4; k++ {
			a, b := u+dd[k], v+dd[k+1]
			if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] != -1 && vis[a][b] != 1 {
				dfs(a, b, sum+1)
			}
		}
		vis[u][v] = 2
	}

	dfs(x, y, 1)

	return res
}
