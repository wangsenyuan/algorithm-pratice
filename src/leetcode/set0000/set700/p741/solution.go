package p741

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, 2)
	dp[0] = make([][]int, n)
	dp[1] = make([][]int, n)

	for i := 0; i < n; i++ {
		dp[0][i] = make([]int, n)
		dp[1][i] = make([]int, n)
	}

	dp[0][0][0] = grid[0][0]
	// dp[1][0][0] = grid[0][0]
	p := 0
	for l := 2; l < n+n; l++ {
		for x1 := 0; x1 < n; x1++ {
			for x2 := 0; x2 < n; x2++ {
				y1 := l - 1 - x1
				y2 := l - 1 - x2
				if y1 < 0 || y1 >= n || y2 < 0 || y2 >= n {
					continue
				}
				if grid[y1][x1] < 0 || grid[y2][x2] < 0 {
					dp[1-p][x1][x2] = -1
					continue
				}
				best := -1
				delta := grid[y1][x1]
				if x1 != x2 {
					delta += grid[y2][x2]
				}
				if x1 > 0 && x2 > 0 && dp[p][x1-1][x2-1] >= 0 {
					best = max(best, dp[p][x1-1][x2-1]+delta)
				}
				if x1 > 0 && y2 > 0 && dp[p][x1-1][x2] >= 0 {
					best = max(best, dp[p][x1-1][x2]+delta)
				}

				if y1 > 0 && x2 > 0 && dp[p][x1][x2-1] >= 0 {
					best = max(best, dp[p][x1][x2-1]+delta)
				}

				if y1 > 0 && y2 > 0 && dp[p][x1][x2] >= 0 {
					best = max(best, dp[p][x1][x2]+delta)
				}
				dp[1-p][x1][x2] = best
			}
		}
		p = 1 - p
	}
	if dp[p][n-1][n-1] < 0 {
		return 0
	}
	return dp[p][n-1][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
