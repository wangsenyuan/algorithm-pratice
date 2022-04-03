package p2222

func numberOfWays(s string) int64 {
	n := len(s)
	// n <= 100000
	dp := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int64, 4)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		// pick x as first times plus one
		dp[x][1]++
		y := 1 - x
		for j := 2; j <= 3; j++ {
			dp[x][j] += dp[y][j-1]
		}
	}

	return dp[0][3] + dp[1][3]
}
