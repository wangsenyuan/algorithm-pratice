package p0091

const INF = 1 << 30

func minCost(costs [][]int) int {
	n := len(costs)

	dp := make([]int, 3)
	dp[0] = costs[0][0]
	dp[1] = costs[0][1]
	dp[2] = costs[0][2]

	fp := make([]int, 3)

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			fp[j] = INF
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				fp[j] = min(fp[j], dp[k]+costs[i][j])
			}
		}
		copy(dp, fp)
	}

	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
