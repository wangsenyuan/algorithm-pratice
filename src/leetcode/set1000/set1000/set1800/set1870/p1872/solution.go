package p1872

func stoneGameVIII(stones []int) int {
	// alice want a - b max
	// bob want a - b min
	n := len(stones)
	pre := make([]int64, n)
	for i := 0; i < n; i++ {
		pre[i] = int64(stones[i])
		if i > 0 {
			pre[i] += pre[i-1]
		}
	}
	dp := make([]int64, n)
	dp[n-1] = pre[n-1]

	for i := n - 2; i > 0; i-- {
		dp[i] = max(dp[i+1], pre[i]-dp[i+1])
	}
	return int(dp[1])
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
