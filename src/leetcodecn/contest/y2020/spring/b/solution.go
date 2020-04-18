package b

func numWays(n int, relation [][]int, k int) int {
	dp := make([]int, n)

	dp2 := make([]int, n)

	dp[0] = 1

	for j := 0; j < k; j++ {
		for _, rel := range relation {
			a, b := rel[0], rel[1]
			dp2[b] += dp[a]
		}

		for i := 0; i < n; i++ {
			dp[i] = dp2[i]
			dp2[i] = 0
		}

	}

	return dp[n-1]
}
