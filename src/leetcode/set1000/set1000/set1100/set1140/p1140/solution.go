package p1140

func stoneGameII(piles []int) int {
	// dp[i][m] = max stones the player can get starting at m
	// dp[0][1] is result

	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, (n+1)/2)
		for j := 0; j < (n+1)/2; j++ {
			dp[i][j] = -1
		}
	}

	dp[n-1][1] = piles[n-1]

	sum := make([]int, n+1)
	sum[n] = 0
	for i := n - 1; i >= 0; i-- {
		sum[i] = piles[i] + sum[i+1]
	}

	var loop func(i int, m int) int
	loop = func(i int, m int) int {
		l := n - i
		if l <= 2*m {
			return sum[i]
		}
		if dp[i][m] >= 0 {
			return dp[i][m]
		}

		var best int
		var s int
		for x := 1; x <= 2*m; x++ {
			s += piles[i+x-1]
			tmp := sum[i+x] - loop(i+x, max(x, m)) + s

			best = max(best, tmp)
		}

		dp[i][m] = best
		return best
	}

	return loop(0, 1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
