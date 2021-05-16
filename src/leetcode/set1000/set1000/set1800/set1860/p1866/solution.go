package p1866

const MOD = 1000000007

func rearrangeSticks(n int, k int) int {
	// dp[i][j] 为有i根，恰好j根可见的数量
	// 增加 第i+1 如果它
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j-1]
			tmp := int64(i-1) * int64(dp[i-1][j]) % MOD
			dp[i][j] += int(tmp)
			dp[i][j] %= MOD
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
