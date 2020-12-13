package p5627

func stoneGameVII(stones []int) int {
	n := len(stones)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
	}

	get := func(i, j int) int {
		if i <= j {
			return dp[i][j]
		}
		return 0
	}

	first := n & 1
	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if (j-i+1)&1 == first {
				// alice turn
				a := sum[j+1] - sum[i+1] + get(i+1, j)
				b := sum[j] - sum[i] + get(i, j-1)
				dp[i][j] = max(a, b)
			} else {
				// bob
				a := get(i+1, j) - (sum[j+1] - sum[i+1])
				b := get(i, j-1) - (sum[j] - sum[i])
				dp[i][j] = min(a, b)
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
