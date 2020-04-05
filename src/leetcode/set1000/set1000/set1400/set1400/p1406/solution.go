package p1406

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)

	get := func(j int) int {
		if j >= n {
			return 0
		}
		return dp[j]
	}
	var sum int
	for i := n - 1; i >= 0; i-- {
		sum += stoneValue[i]
		dp[i] = sum - get(i+1)
		dp[i] = max(dp[i], sum-get(i+2))
		dp[i] = max(dp[i], sum-get(i+3))
	}

	took := dp[0]

	if took > sum-took {
		return "Alice"
	}
	if took < sum-took {
		return "Bob"
	}

	return "Tie"
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func stoneGameIII1(stoneValue []int) string {
	n := len(stoneValue)

	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + stoneValue[i]
	}

	dp := make([]int, n+1)
	vis := make([]bool, n+1)
	vis[n] = true
	dp[n] = 0

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		if vis[i] {
			return dp[i]
		}
		vis[i] = true

		dp[i] = math.MinInt32
		for j := 1; j <= 3 && i+j <= n; j++ {
			tmp := dfs(i + j)
			cur := sum[i] - tmp
			dp[i] = max(dp[i], cur)
		}

		return dp[i]
	}

	took := dfs(0)

	if took > sum[0]-took {
		return "Alice"
	}
	if took < sum[0]-took {
		return "Bob"
	}

	return "Tie"
}
