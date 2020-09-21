package p1595

import "math"

const MOD = 1000000007

func maxProductPath(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int64, 2)
		}
	}
	dp[0][0][0] = int64(grid[0][0])
	dp[0][0][1] = int64(grid[0][0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid[i][j] == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = 0
				continue
			}
			dp[i][j][0] = math.MinInt32
			dp[i][j][1] = 1 << 20
			if i > 0 {
				if grid[i][j] > 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0]*int64(grid[i][j]))
					dp[i][j][1] = min(dp[i][j][1], dp[i-1][j][1]*int64(grid[i][j]))
				} else {
					dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][1]*int64(grid[i][j]))
					dp[i][j][1] = min(dp[i][j][1], dp[i-1][j][0]*int64(grid[i][j]))
				}
			}
			if j > 0 {
				if grid[i][j] > 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][0]*int64(grid[i][j]))
					dp[i][j][1] = min(dp[i][j][1], dp[i][j-1][1]*int64(grid[i][j]))
				} else {
					dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][1]*int64(grid[i][j]))
					dp[i][j][1] = min(dp[i][j][1], dp[i][j-1][0]*int64(grid[i][j]))
				}
			}
		}
	}

	if dp[m-1][n-1][0] < 0 {
		return -1
	}

	return int(dp[m-1][n-1][0] % MOD)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
