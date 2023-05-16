package p1335

func minDifficulty(jobDifficulty []int, d int) int {
	// result = sum(day difficulty)
	// day difficulty = max(job difficulty in that day)
	// 任务肯定是按照顺序完成的
	// 假设到第i天是完成了前j项任务的总的最小数量
	// dp[i][j] = dp[i-1][k] + max(diff[k+1...j])
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	const inf = 1 << 30

	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = inf
	}

	dp[0] = 0

	fp := make([]int, n+1)

	for i := 1; i <= d; i++ {

		for j := 0; j <= n; j++ {
			fp[j] = inf
		}

		for j := i; j <= n; j++ {
			x := jobDifficulty[j-1]
			for k := j - 1; k >= 0; k-- {
				fp[j] = min(fp[j], dp[k]+x)
				if k > 0 {
					x = max(x, jobDifficulty[k-1])
				}
			}
		}
		copy(dp, fp)
	}

	return dp[n]
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
