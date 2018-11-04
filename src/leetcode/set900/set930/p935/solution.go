package p935

const MOD = 1e9 + 7

func knightDialer(N int) int {
	if N == 1 {
		return 10
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}

	var p int
	for i := 1; i < N; i++ {
		q := 1 - p
		dp[q][0] = modAdd(dp[p][4], dp[p][6])
		dp[q][1] = modAdd(dp[p][6], dp[p][8])
		dp[q][2] = modAdd(dp[p][7], dp[p][9])
		dp[q][3] = modAdd(dp[p][4], dp[p][8])
		dp[q][4] = modAdd(dp[p][3], dp[p][9])
		dp[q][4] = modAdd(dp[q][4], dp[p][0])
		dp[q][5] = 0
		dp[q][6] = modAdd(dp[p][1], dp[p][7])
		dp[q][6] = modAdd(dp[q][6], dp[p][0])
		dp[q][7] = modAdd(dp[p][2], dp[p][6])
		dp[q][8] = modAdd(dp[p][1], dp[p][3])
		dp[q][9] = modAdd(dp[p][2], dp[p][4])
		p = q
	}

	var res int
	for i := 0; i < 10; i++ {
		res = modAdd(res, dp[p][i])
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
