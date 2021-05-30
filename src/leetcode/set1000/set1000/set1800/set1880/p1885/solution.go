package p1885

import "math"

const EPS = 1e-8

func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)

	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = float64(hoursBefore + 1)
		}
	}
	S := float64(speed)
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j < i {
				dp[i][j] = math.Ceil(dp[i-1][j] + float64(dist[i-1])/S - EPS)
			}
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+float64(dist[i-1])/S)
			}
		}
	}

	for j := 0; j <= n; j++ {
		if dp[n][j] <= float64(hoursBefore)+EPS {
			return j
		}
	}
	return -1
}

func min(a, b float64) float64 {
	return math.Min(a, b)
}
