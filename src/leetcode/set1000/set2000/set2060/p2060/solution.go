package p2060

func possiblyEquals(s1 string, s2 string) bool {
	// dp[i][j] = true
	m := len(s1)
	n := len(s2)

	dp := make([][][][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([][]int, 2)
			for w := 0; w < 2; w++ {
				dp[i][j][w] = make([]int, 1000)
				for r := 0; r < 1000; r++ {
					dp[i][j][w][r] = -1
				}
			}
		}
	}

	var fn func(i, j, which, rest int) int
	fn = func(i, j, which, rest int) int {
		if dp[i][j][which][rest] >= 0 {
			return dp[i][j][which][rest]
		}
		dp[i][j][which][rest] = 0
		if which == 0 {
			if j == n {
				if i == m && rest == 0 {
					dp[i][j][which][rest] = 1
				}
			} else if isAlpha(s2[j]) {
				if rest == 0 && i < m && isAlpha(s1[i]) {
					if s1[i] == s2[j] && fn(i+1, j+1, 0, 0) > 0 {
						dp[i][j][which][rest] = 1
					}
				} else if rest >= 1 && fn(i, j+1, 0, rest-1) > 0 || rest == 0 && fn(i, j+1, 1, 1) > 0 {
					dp[i][j][which][rest] = 1
				}
			} else {
				var x, k = 0, j
				for k < n && !isAlpha(s2[k]) {
					x = x*10 + int(s2[k]-'0')
					if rest >= x && fn(i, k+1, 0, rest-x) > 0 || rest < x && fn(i, k+1, 1, x-rest) > 0 {
						dp[i][j][which][rest] = 1
						break
					}
					k++
				}
			}
		} else {
			if i == m {
				if j == n && rest == 0 {
					dp[i][j][which][rest] = 1
				}
			} else if isAlpha(s1[i]) {
				if rest == 0 && j < n && isAlpha(s2[j]) {
					if s1[i] == s2[j] && fn(i+1, j+1, 1, 0) > 0 {
						dp[i][j][which][rest] = 1
					}
				} else if rest >= 1 && fn(i+1, j, 1, rest-1) > 0 || rest == 0 && fn(i+1, j, 0, 1) > 0 {
					dp[i][j][which][rest] = 1
				}
			} else {
				var x, k = 0, i
				for k < m && !isAlpha(s1[k]) {
					x = x*10 + int(s1[k]-'0')
					if rest >= x && fn(k+1, j, 1, rest-x) > 0 || rest < x && fn(k+1, j, 0, x-rest) > 0 {
						dp[i][j][which][rest] = 1
						break
					}
					k++
				}
			}
		}

		return dp[i][j][which][rest]
	}

	res := fn(0, 0, 0, 0)

	return res > 0
}

func isAlpha(x byte) bool {
	return x >= 'a' && x <= 'z'
}
