package p1230

func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([]float64, target+1)
	fp := make([]float64, target+1)

	dp[0] = 1.0

	for i := 0; i < n; i++ {
		p := prob[i]

		fp[0] = dp[0] * (1 - p)

		x := min(target, i+1)

		for j := 1; j <= x; j++ {
			fp[j] = dp[j]*(1-p) + dp[j-1]*p
		}

		copy(dp, fp)
	}

	return dp[target]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
