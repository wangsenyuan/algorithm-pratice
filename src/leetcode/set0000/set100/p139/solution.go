package p139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	ws := make(map[string]bool)

	for _, w := range wordDict {
		ws[w] = true
	}

	dp := make([]bool, n + 1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := i - 1; !dp[i] && j >= 0; j-- {
			dp[i] =  ws[s[j:i]] && dp[j]
		} 
	}

	return dp[n]
}

func wordBreak1(s string, wordDict []string) bool {
	ws := make(map[string]bool)

	for _, w := range wordDict {
		ws[w] = true
	}

	dp := make([]int, len(s))


	var dfs func(i int) bool
	
	dfs = func(i int) bool {
		if i == len(s) {
			return true
		}

		if dp[i] != 0 {
			return dp[i] > 0
		}

		for j := i + 1; j <= len(s); j++ {
			if ws[s[i:j]] && dfs(j) {
				dp[i] = 1
				return true
			}
		}
		dp[i] = -1
		return false
	}

	return dfs(0)
}