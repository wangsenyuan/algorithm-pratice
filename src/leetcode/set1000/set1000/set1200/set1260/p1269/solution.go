package p1269

const MOD = 1000000007

func numWays(steps int, arrLen int) int {
	// points far away from point 0, will not come base
	n := min(steps, arrLen)
	// n = (n + 1) / 2
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	fp := make([]int, n)
	dp[0] = 1

	for i := 0; i < steps; i++ {
		fp[0] = dp[0] + dp[1]
		if fp[0] >= MOD {
			fp[0] -= MOD
		}
		for j := 1; j < n-1; j++ {
			fp[j] = dp[j] + dp[j-1]
			if fp[j] >= MOD {
				fp[j] -= MOD
			}
			fp[j] += dp[j+1]
			if fp[j] >= MOD {
				fp[j] -= MOD
			}
		}
		fp[n-1] = dp[n-1] + dp[n-2]

		if fp[n-1] >= MOD {
			fp[n-1] -= MOD
		}

		for j := 0; j < n; j++ {
			dp[j] = fp[j]
			fp[j] = 0
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
