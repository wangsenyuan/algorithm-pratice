package p3258

func maxEnergyBoost(a []int, b []int) int64 {
	n := len(a)
	// dp[i] = 到i小时的最大能量
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[1][0] = a[0]
	dp[1][1] = b[0]
	for i := 1; i <= n; i++ {
		if i == 1 {
			// 相同的
			dp[i][0] = dp[i-1][0] + a[i-1]
			dp[i][1] = dp[i-1][1] + b[i-1]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-2][1]) + a[i-1]
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]) + b[i-1]
		}
	}

	return int64(max(dp[n][0], dp[n][1]))
}
