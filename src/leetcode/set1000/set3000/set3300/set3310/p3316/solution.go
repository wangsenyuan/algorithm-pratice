package p3316

const inf = 1 << 60

func maxRemovals(source string, pattern string, targetIndices []int) int {
	// 就是两个p[i], p[i+1] 中间在target 中的都可以删除
	// 但是问题是 p[i], p[i+1] 有多个匹配
	m := len(pattern)
	n := len(source)
	value := make([]int, n)
	for _, i := range targetIndices {
		value[i] = 1
	}

	dp := make([]int, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = -inf
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			// dp[i][j] = dp[i-1][j] + value[i]
			tmp := dp[j] + value[i]
			if j > 0 && source[i] == pattern[j-1] {
				// dp[i][j] = dp[i-1][j-1]
				tmp = max(tmp, dp[j-1])
			}
			dp[j] = tmp
		}
	}
	return dp[m]
}
