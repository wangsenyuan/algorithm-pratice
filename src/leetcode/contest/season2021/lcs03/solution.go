package lcs03

var dd = []int{-1, 0, 1, 0, -1}

func largestArea(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}

	var dfs func(i, j int, x byte) int

	dfs = func(i, j int, x byte) int {
		res := 1
		vis[i][j]++
		for k := 0; k < 4; k++ {
			u, v := i+dd[k], j+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && vis[u][v] == 0 && grid[u][v] == x {
				res += dfs(u, v, x)
			}
		}
		return res
	}

	var dfs2 func(i, j int, x byte) bool

	dfs2 = func(i, j int, x byte) bool {
		vis[i][j]++
		for k := 0; k < 4; k++ {
			u, v := i+dd[k], j+dd[k+1]
			if u < 0 || u == m || v < 0 || v == n || grid[u][v] == '0' {
				return true
			}
			if vis[u][v] == 1 && grid[u][v] == x && dfs2(u, v, x) {
				return true
			}
		}

		return false
	}

	var best int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] > 0 {
				continue
			}
			if grid[i][j] == '0' {
				continue
			}
			cnt := dfs(i, j, grid[i][j])
			if cnt <= best {
				continue
			}
			if !dfs2(i, j, grid[i][j]) {
				best = cnt
			}
		}
	}
	return best
}
