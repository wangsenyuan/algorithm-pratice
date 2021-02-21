package p1769
func maximumScore(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, m+1)
	}

	// dp[i][j] = 生成 长度为i的结果数组，且左边的index为j时的最大值
	dp[0][0] = 0
	var p int
	for l := 1; l <= m; l++ {
		q := p ^ 1
		for j := 0; j < m+1; j++ {
			dp[q][j] = -(1 << 30)
		}
		for i := 0; i < l; i++ {
			// dp[l][i] = ?
			dp[q][i] = dp[p][i] + nums[n-(l-i)]*multipliers[l-1]
			if i > 0 {
				dp[q][i] = max(dp[q][i], dp[p][i-1]+nums[i-1]*multipliers[l-1])
			}
		}
		dp[q][l] = dp[p][l-1] + nums[l-1]*multipliers[l-1]
		p = q
	}

	var best = -(1 << 30)

	for i := 0; i < m+1; i++ {
		best = max(best, dp[p][i])
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}