package main

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			dp[i+1][j+1] = dp[i][j] + 1
			if dp[i][j+1]+1 < dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i][j+1] + 1
			}

			if dp[i+1][j]+1 < dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i+1][j] + 1
			}
		}
	}

	return dp[m][n]
}
