package p1402

import "sort"

const INF = 1 << 30
const N_INF = -INF

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	n := len(satisfaction)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = N_INF
		}
	}

	for i := 1; i <= n; i++ {
		// prepare i
		for k := 0; k < i; k++ {
			dp[i][k] = max(dp[i][k], dp[i-1][k])
			dp[i][k+1] = max(dp[i][k+1], dp[i-1][k]+satisfaction[i-1]*(k+1))
		}
	}

	var res int

	for k := 0; k <= n; k++ {
		res = max(res, dp[n][k])
	}

	return res
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
