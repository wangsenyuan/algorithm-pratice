package p5555

func countVowelStrings(n int) int {
	dp := make([]int, 5)
	dp[0] = 1

	for i := 0; i < n; i++ {
		dp[4] += dp[3] + dp[2] + dp[1] + dp[0]
		dp[3] += dp[2] + dp[1] + dp[0]
		dp[2] += dp[1] + dp[0]
		dp[1] += dp[0]
	}

	return dp[4] + dp[3] + dp[2] + dp[1] + dp[0]
}
