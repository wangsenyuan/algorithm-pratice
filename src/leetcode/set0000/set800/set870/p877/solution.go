package p877

func stoneGame(piles []int) bool {
	//dp[i][j][0] = alice got max when array between [i, j] inclusive
	//dp[i][j][0] = sum[i][j] - dp[i+1][j][1] + A[i] or sum[i][j] - dp[i][j-1][1] + A[j]

	n := len(piles)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + piles[i]
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -1
			dp[i][j][1] = -1
		}
	}

	var dfs func(i, j int, turn int) int

	dfs = func(i, j, turn int) int {
		if dp[i][j][turn] < 0 {
			if i == j {
				dp[i][j][turn] = piles[i]
			} else {
				a := dfs(i+1, j, 1-turn)
				a = sum[j+1] - sum[i] - a + piles[i]
				b := dfs(i, j-1, 1-turn)
				b = sum[j+1] - sum[i] - b + piles[j]
				dp[i][j][turn] = max(a, b)
			}
		}

		return dp[i][j][turn]
	}

	return dfs(0, n-1, 0)*2 > sum[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
