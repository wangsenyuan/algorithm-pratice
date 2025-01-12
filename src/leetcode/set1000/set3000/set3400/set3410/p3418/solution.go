package p3418

const inf = 1 << 60

func maximumAmount(coins [][]int) int {
	n := len(coins)
	m := len(coins[0])
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				dp[i][j][k] = -inf
			}
		}
	}

	dp[0][0][0] = coins[0][0]
	if coins[0][0] < 0 {
		dp[0][0][1] = 0
	}

	get := func(i int, j int, k int) int {
		if i < 0 || j < 0 {
			return -inf
		}
		return dp[i][j][k]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 3; k++ {
				dp[i][j][k] = max(dp[i][j][k], get(i-1, j, k)+coins[i][j], get(i, j-1, k)+coins[i][j])

				if coins[i][j] < 0 && k+1 < 3 {
					dp[i][j][k+1] = max(dp[i][j][k+1], get(i-1, j, k), get(i, j-1, k))
				}
			}
		}
	}

	return max(dp[n-1][m-1][0], dp[n-1][m-1][1], dp[n-1][m-1][2])
}
