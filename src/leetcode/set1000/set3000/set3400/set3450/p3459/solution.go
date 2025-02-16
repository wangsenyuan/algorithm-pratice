package p3459

func lenOfVDiagonal(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][]int, n)
	fp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		fp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 4)
			fp[i][j] = make([]int, 4)
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j]%2 == 0 {
				dp[i][j][0] = 1
				if i+1 < n && j+1 < m && grid[i+1][j+1]+grid[i][j] == 2 {
					dp[i][j][0] = 1 + dp[i+1][j+1][0]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if grid[i][j]%2 == 0 {
				dp[i][j][1] = 1
				if i+1 < n && j-1 >= 0 && grid[i+1][j-1]+grid[i][j] == 2 {
					dp[i][j][1] = 1 + dp[i+1][j-1][1]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j]%2 == 0 {
				dp[i][j][2] = 1
				if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1]+grid[i][j] == 2 {
					dp[i][j][2] = 1 + dp[i-1][j-1][2]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j]%2 == 0 {
				dp[i][j][3] = 1
				if i-1 >= 0 && j+1 < m && grid[i-1][j+1]+grid[i][j] == 2 {
					dp[i][j][3] = 1 + dp[i-1][j+1][3]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			// fp[i][j][0] = dp[i+x][j+x][1] + x
			if grid[i][j]%2 == 0 {
				fp[i][j][0] = dp[i][j][1]
				if dp[i][j][0] > 1 {
					fp[i][j][0] = max(fp[i][j][0], fp[i+1][j+1][0]+1)
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if grid[i][j]%2 == 0 {
				fp[i][j][1] = dp[i][j][2]
				if dp[i][j][1] > 1 {
					fp[i][j][1] = max(fp[i][j][1], fp[i+1][j-1][1]+1)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j]%2 == 0 {
				fp[i][j][2] = dp[i][j][3]
				if dp[i][j][2] > 1 {
					fp[i][j][2] = max(fp[i][j][2], fp[i-1][j-1][2]+1)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j]%2 == 0 {
				fp[i][j][3] = dp[i][j][0]
				if dp[i][j][3] > 1 {
					fp[i][j][3] = max(fp[i][j][3], fp[i-1][j+1][3]+1)
				}
			}
		}
	}

	best := -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				best = max(best, 0)
				if i+1 < n && j+1 < m && grid[i+1][j+1] == 2 {
					best = max(best, dp[i+1][j+1][0], fp[i+1][j+1][0])
				}
				if i+1 < n && j-1 >= 0 && grid[i+1][j-1] == 2 {
					best = max(best, dp[i+1][j-1][1], fp[i+1][j-1][1])
				}
				if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == 2 {
					best = max(best, dp[i-1][j-1][2], fp[i-1][j-1][2])
				}
				if i-1 >= 0 && j+1 < m && grid[i-1][j+1] == 2 {
					best = max(best, dp[i-1][j+1][3], fp[i-1][j+1][3])
				}
			}
		}
	}

	return best + 1
}
