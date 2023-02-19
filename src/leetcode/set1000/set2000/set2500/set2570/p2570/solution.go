package p2570

func findTheString(lcp [][]int) string {
	n := len(lcp)
	// lcp[i][j] = 最长common prefix string in word[i..] and word[j...]
	// lcp[i][j] = l
	// word[i...i+l] = word[j...j+l]
	// => word[i+l] != word[j+l]
	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			l := lcp[i][j]
			if i+l > n || j+l > n || i == j && lcp[i][j] != n-i {
				return ""
			}

			if i+l < n && j+l < n {
				conn[i+l][j+l] = 1
			}
		}
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		var flag int
		for j := i - 1; j >= 0; j-- {
			if conn[j][i] == 1 {
				flag |= (1 << int(buf[j]-'a'))
			}
		}
		buf[i] = ' '
		for x := 0; x < 26; x++ {
			if (flag>>x)&1 == 0 {
				buf[i] = byte(x + 'a')
				break
			}
		}
		if buf[i] == ' ' {
			return ""
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if buf[i] == buf[j] {
				dp[i][j]++
				if i+1 < n && j+1 < n {
					dp[i][j] += dp[i+1][j+1]
				}
			}
			if lcp[i][j] != dp[i][j] {
				return ""
			}
		}
	}

	return string(buf)
}
