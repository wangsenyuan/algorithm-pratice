package p1906

var dd = []int{-1, 0, 1, 0, -1}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])

	vis := make([][]bool, m)

	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		vis[x][y] = true
		covered := true
		if grid1[x][y] == 0 {
			covered = false
		}

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && grid2[u][v] == 1 && !vis[u][v] {
				tmp := dfs(u, v)
				if tmp < 0 {
					covered = false
				}
			}
		}

		if !covered {
			return -1
		}
		return 1
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && !vis[i][j] {
				tmp := dfs(i, j)
				if tmp == 1 {
					res++
				}
			}
		}
	}
	return res
}
