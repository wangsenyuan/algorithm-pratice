package p329

var dd = []int{-1, 0, 1, 0, -1}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		if dp[x][y] > 0 {
			return dp[x][y]
		}
		dp[x][y] = 1
		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u >= 0 && u < m && v >= 0 && v < n && matrix[u][v] > matrix[x][y] {
				dp[x][y] = max(dp[x][y], dfs(u, v)+1)
			}
		}
		return dp[x][y]
	}
	var best = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			best = max(best, dfs(i, j))
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
