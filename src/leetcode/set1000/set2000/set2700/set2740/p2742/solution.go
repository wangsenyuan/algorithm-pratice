package p2742

const INF = 1 << 30

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	// 应该尽可能的让免费的油漆工工作
	// dp[time] = 在时间x内完成的最小的cost
	// 如果 n - time <= time, good
	// time <= n
	// 如果付费工工作了时间t，那么免费工，可以完成t个任务，那么这时候，把最cost的任务都交给免费工即可，
	// 此时，只要剩余的任务的time sum >= t即可

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(i int, j int) int

	dfs = func(i int, j int) int {
		if j-n > i {
			return 0
		}
		if i < 0 {
			return INF
		}
		if dp[i][j] == -1 {
			dp[i][j] = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		}
		return dp[i][j]
	}

	return dfs(n-1, n)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
