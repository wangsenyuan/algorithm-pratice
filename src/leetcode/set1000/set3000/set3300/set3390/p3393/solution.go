package p3393

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}
func countPathsWithXorValue(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 16)
		}
	}

	dp[0][0][grid[0][0]] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 {
				for x := 0; x < 16; x++ {
					dp[i][j][x^grid[i][j]] = add(dp[i][j][x^grid[i][j]], dp[i-1][j][x])
				}
			}
			if j > 0 {
				for x := 0; x < 16; x++ {
					dp[i][j][x^grid[i][j]] = add(dp[i][j][x^grid[i][j]], dp[i][j-1][x])
				}
			}
		}
	}
	return dp[n-1][m-1][k]
}
