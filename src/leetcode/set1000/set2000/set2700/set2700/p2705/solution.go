package p2705

func minExtraChar(s string, dictionary []string) int {
	words := make(map[string]bool)

	for _, w := range dictionary {
		words[w] = true
	}

	n := len(s)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if words[s[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
