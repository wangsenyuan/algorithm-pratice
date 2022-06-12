package p2304

const INF = 1 << 30

func minPathCost(grid [][]int, moveCost [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}

	for j := 0; j < n; j++ {
		dp[0][j] = grid[0][j]
	}

	for r := 1; r < m; r++ {
		for i := 0; i < n; i++ {
			// if prev at col i
			for j := 0; j < n; j++ {
				// move to col j
				dp[r][j] = min(dp[r][j], dp[r-1][i]+moveCost[grid[r-1][i]][j]+grid[r][j])
			}
		}
	}

	res := INF

	for j := 0; j < n; j++ {
		res = min(res, dp[m-1][j])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
