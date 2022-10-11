package p2435

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func numberOfPaths(grid [][]int, k int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k)
		}
	}

	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for x := 0; x < k; x++ {
				// go right
				if j+1 < n {
					nx := (x + grid[i][j+1]) % k
					dp[i][j+1][nx] = modAdd(dp[i][j+1][nx], dp[i][j][x])
				}
				if i+1 < m {
					nx := (x + grid[i+1][j]) % k
					dp[i+1][j][nx] = modAdd(dp[i+1][j][nx], dp[i][j][x])
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}
