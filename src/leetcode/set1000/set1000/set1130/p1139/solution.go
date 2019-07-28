package p1139

func largest1BorderedSquare(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	m := len(grid[0])
	if m == 0 {
		return 0
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = 0
			} else {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				if i > 0 {
					dp[i][j][0] += dp[i-1][j][0]
				}
				if j > 0 {
					dp[i][j][1] += dp[i][j-1][1]
				}
			}
		}
	}

	w := min(m, n)

	for w >= 1 {
		for i := w - 1; i < n; i++ {
			for j := w - 1; j < m; j++ {
				if dp[i][j][0] >= w && dp[i][j][1] >= w {
					u := i - w + 1
					v := j - w + 1

					if dp[u][j][1] >= w && dp[i][v][0] >= w {
						return w * w
					}
				}
			}
		}
		w--
	}

	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
