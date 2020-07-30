package p343

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for x := 2; x <= n; x++ {
		for y := 1; y < x; y++ {
			z := x - y
			dp[x] = max(dp[x], dp[y]*dp[z])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
