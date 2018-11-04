package p935

const MOD = 1e9 + 7

func knightDialer(N int) int {
	dp := make([][]int, N+5)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < N; i++ {
		dp[i][0] = modAdd(dp[i-1][4], dp[i-1][6])
		dp[i][1] = modAdd(dp[i-1][6], dp[i-1][8])
		dp[i][2] = modAdd(dp[i-1][7], dp[i-1][9])
		dp[i][3] = modAdd(dp[i-1][4], dp[i-1][8])
		dp[i][4] = modAdd(dp[i-1][3], dp[i-1][9])
		dp[i][4] = modAdd(dp[i][4], dp[i-1][0])
		dp[i][5] = 0
		dp[i][6] = modAdd(dp[i-1][1], dp[i-1][7])
		dp[i][6] = modAdd(dp[i][6], dp[i-1][0])
		dp[i][7] = modAdd(dp[i-1][2], dp[i-1][6])
		dp[i][8] = modAdd(dp[i-1][1], dp[i-1][3])
		dp[i][9] = modAdd(dp[i-1][2], dp[i-1][4])
	}

	var res int
	for i := 0; i < 10; i++ {
		res = modAdd(res, dp[N-1][i])
	}
	return res
}

func modAdd(a, b int) int {
	c := a + b
	if c >= MOD {
		c -= MOD
	}
	return c
}
