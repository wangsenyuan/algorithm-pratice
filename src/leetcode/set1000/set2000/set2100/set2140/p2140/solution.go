package p2140

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	// dp[i] = 从问题i开始解决的最大值
	// dp[j] = max(dp[j+1], points[j] + dp[j+pow[j]+1])

	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		cur := questions[i]
		j := min(n, i+cur[1]+1)
		dp[i] = max2(dp[i+1], dp[j]+int64(cur[0]))
	}

	return dp[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
