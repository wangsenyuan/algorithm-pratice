package p1873

const INF = 1 << 30

func twoEggDrop(n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	dp[1][0] = 0

	for j := 1; j <= n; j++ {
		dp[0][j] = j
		for k := 1; k <= j; k++ {
			dp[1][j] = min(dp[1][j], max(dp[0][k-1], dp[1][j-k])+1)
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
