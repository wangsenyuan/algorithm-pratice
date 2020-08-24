package p1563

const N_INF = -(1 << 20)

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = N_INF
		}
		dp[i][i] = 0
	}

	ps := make([]int, n+1)

	for i := 0; i < n; i++ {
		ps[i+1] = ps[i] + stoneValue[i]
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			//dp[i][j] = max(dp[i][k] + ps[k+1] - ps[i], dp[k][j] + ps[j+1] - ps[k])
			for k := i; k < j; k++ {
				if ps[k+1]-ps[i] < ps[j+1]-ps[k+1] {
					dp[i][j] = max(dp[i][j], dp[i][k]+ps[k+1]-ps[i])
				} else if ps[k+1]-ps[i] > ps[j+1]-ps[k+1] {
					dp[i][j] = max(dp[i][j], dp[k+1][j]+ps[j+1]-ps[k+1])
				} else {
					dp[i][j] = max(dp[i][j], max(dp[i][k], dp[k+1][j])+ps[j+1]-ps[k+1])
				}
			}
		}
	}

	return dp[0][n-1]
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
