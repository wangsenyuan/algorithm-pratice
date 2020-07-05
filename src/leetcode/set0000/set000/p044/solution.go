package p044

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	if n == 0 {
		return m == 0
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	if p[0] == '*' {
		for i := 0; i <= m; i++ {
			dp[i][1] = true
		}
	}

	for j := 0; j < n; j++ {
		if p[j] == '*' {
			for i := 0; i <= m; i++ {
				dp[i][j+1] = dp[i][j]
			}
		}
		for i := 0; i < m; i++ {
			if p[j] == '?' || s[i] == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				// empty
				dp[i+1][j+1] = dp[i+1][j]
				for ii := i; ii >= 0 && !dp[i+1][j+1]; ii-- {
					dp[i+1][j+1] = dp[ii][j]
				}
			}
		}
	}

	return dp[m][n]
}
