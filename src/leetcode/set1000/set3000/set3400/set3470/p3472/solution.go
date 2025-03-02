package p3472

func longestPalindromicSubsequence(s string, k int) int {
	n := len(s)
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, k+1)
		}
		dp[i][i][0] = 1
	}

	for r := range n {
		for l := r - 1; l >= 0; l-- {
			for x := 0; x <= k; x++ {
				dp[l][r][x] = max(dp[l][r][x], dp[l+1][r][x], dp[l][r-1][x])
				u := get(s[l])
				v := get(s[r])
				if u > v {
					u, v = v, u
				}
				// u - v
				diff := min(v-u, u+26-v)
				if x >= diff {
					dp[l][r][x] = max(dp[l][r][x], dp[l+1][r-1][x-diff]+2)
				}
				if x > 0 {
					dp[l][r][x] = max(dp[l][r][x], dp[l][r][x-1])
				}
			}
		}
	}
	return max(1, dp[0][n-1][k])
}

func get(c byte) int {
	return int(c - 'a')
}
func abs(num int) int {
	return max(num, -num)
}
