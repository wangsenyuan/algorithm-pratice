package p2472

func maxPalindromes(s string, k int) int {
	n := len(s)
	// n <= 2000
	pal := make([][]bool, n)
	// pal[i][j] = s[i:j+1] is palindrome or not
	for i := 0; i < n; i++ {
		pal[i] = make([]bool, n)
		pal[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if i+1 <= j-1 {
					pal[i][j] = pal[i+1][j-1]
				} else {
					pal[i][j] = true
				}
			} else {
				pal[i][j] = false
			}
		}
	}
	// dp[i] = max num until i
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		for j := i; j >= 0; j-- {
			if pal[j][i] && (i-j+1) >= k {
				if j == 0 {
					dp[i] = max(dp[i], 1)
				} else {
					dp[i] = max(dp[i], dp[j-1]+1)
				}
			}
		}
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
