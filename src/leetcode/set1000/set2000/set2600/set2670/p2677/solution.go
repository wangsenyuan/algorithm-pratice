package p2677

func maxMoves(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 1; j < m; j++ {
			dp[i][j] = -m
		}
	}

	var res int
	for j := 1; j < m; j++ {
		for i := 0; i < n; i++ {
			for k := -1; k <= 1; k++ {
				if i+k >= 0 && i+k < n && grid[i][j] > grid[i+k][j-1] {
					dp[i][j] = max(dp[i][j], dp[i+k][j-1]+1)
				}
			}

			res = max(res, dp[i][j])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
