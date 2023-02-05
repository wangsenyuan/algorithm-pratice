package p2560

func isPossibleToCutPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	cnt1 := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && dp[i-1][j] == 1 && grid[i][j] == 1 {
				dp[i][j] = 1
			}
			if j > 0 && dp[i][j-1] == 1 && grid[i][j] == 1 {
				dp[i][j] = 1
			}
			cnt1[i+j] += dp[i][j]
		}
	}

	for k, v := range cnt1 {
		if k == 0 || k == m-1+n-1 {
			continue
		}
		if v <= 1 {
			return true
		}
	}

	return false
}
