package lcp47

const MOD = 1000000007

func securityCheck(capacities []int, k int) int {
	// k <= sum(cap)
	// dp[i][x][j]
	// 表示在第i个检查室时，第k个人位于第j位时，且(x = 0 为que，1 为stack时的数量)
	// dp[i+1][nx][nj] += dp[i][x][j]
	n := len(capacities)

	dp := make([]int, k+1)
	dp[k] = 1

	for i := 0; i < n; i++ {
		rem := capacities[i] - 1
		for j := rem; j <= k; j++ {
			dp[j-rem] += dp[j]
			if dp[j-rem] >= MOD {
				dp[j-rem] -= MOD
			}
		}
	}

	return dp[0]
}
