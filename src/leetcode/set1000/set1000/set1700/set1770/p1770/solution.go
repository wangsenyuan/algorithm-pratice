package p1770

func longestPalindrome(word1 string, word2 string) int {
	// sub1 + sub2 is palindrome
	// word1 + word2 找出最长的回文串
	dp := palindrome(word1 + word2)

	m := len(word1)
	n := len(word2)

	var best int

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				// candidate position
				ii := i + 1
				jj := m + j - 1
				tmp := m - i + j + 1 - dp[ii][jj]
				best = max(best, tmp)
			}
		}
	}

	return best
}

func palindrome(word string) [][]int {
	n := len(word)
	//dp[i][j]使得word[i:j]为回文串删除字符的最小值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = n
		}
	}

	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if word[i] == word[j] {
				// can keep i & j
				dp[i][j] = dp[i+1][j-1]
			} else {
				// remove both
				dp[i][j] = dp[i+1][j-1] + 2
				dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	return dp
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
