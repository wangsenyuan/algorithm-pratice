package p2402

func longestSubsequence(s string, k int) int {
	n := len(s)

	// sub_seq(s[i:...]) <= k
	var best int
	var sum int
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			sum++
		} else {
			// s[i] == '1'
			// at most length(k)
			a := sum
			var cur int
			for j := i; j < n && cur < k; j++ {
				cur = (cur << 1) | int(s[j]-'0')
				a++
				if cur <= k {
					best = max(best, a)
				}
			}
		}
	}

	return max(best, sum)
}

func longestSubsequence1(s string, k int) int {
	// len(s) <= 1000
	var arr []byte

	for k > 0 {
		x := byte((k & 1) + '0')
		arr = append(arr, x)
		k >>= 1
	}
	reverse(arr)
	m := len(s)
	n := len(arr)
	dp := make([][][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i <= m; i++ {
		dp[i][0][1] = -(1 << 30)
	}

	// dp[i][j][0] = till i, j, get equal sub sequence length
	// dp[i][j][1] = till i, j, get sub sequence <= arr[:j] length
	for j := 0; j <= n; j++ {
		for i := 1; i <= m; i++ {
			x := int(s[i-1] - '0')
			if j == 0 {
				if x == 0 {
					dp[i][j][0] = 1 + dp[i-1][j][0]
				} else {
					dp[i][j][0] = dp[i-1][j][0]
				}
			} else {
				// j > 0
				y := int(arr[j-1] - '0')
				if x == y {
					// keep i or skip i
					dp[i][j][0] = max(dp[i-1][j-1][0]+1, dp[i-1][j][0])
					dp[i][j][1] = max(dp[i-1][j-1][1]+1, dp[i-1][j][1])
				} else if x < y {
					// only skip it
					dp[i][j][0] = dp[i-1][j][0]
					// skip i or keep i
					dp[i][j][1] = max(dp[i-1][j][1], max(dp[i-1][j-1][1], dp[i-1][j-1][0])+1)
				} else {
					// x > y
					dp[i][j][0] = dp[i-1][j][0]
					dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][1]+1)
				}
			}
		}
	}

	return max(dp[m][n][0], dp[m][n][1])
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
