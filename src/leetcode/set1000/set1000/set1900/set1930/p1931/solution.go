package p1931

const INF = 1 << 30

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, maxTime+1)
		for j := 0; j <= maxTime; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = passingFees[0]

	for t := 1; t <= maxTime; t++ {
		for _, cur := range edges {
			x, y, w := cur[0], cur[1], cur[2]
			if w <= t {
				dp[x][t] = min(dp[x][t], dp[y][t-w]+passingFees[x])
				dp[y][t] = min(dp[y][t], dp[x][t-w]+passingFees[y])
			}
		}
	}

	best := INF

	for t := 1; t <= maxTime; t++ {
		best = min(best, dp[n-1][t])
	}
	if best == INF {
		return -1
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
