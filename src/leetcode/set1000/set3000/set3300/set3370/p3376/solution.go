package p3376

const inf = 1 << 60

func findMinimumTime(strength []int, K int) int {
	n := len(strength)
	// n <= 8
	// dp[state] 表示处理完state表示的，所需的最小时间
	N := 1 << n

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = inf
	}

	dp[0] = 0

	for state := 0; state < N; state++ {
		var cnt int
		for i := 0; i < n; i++ {
			cnt += (state >> i) & 1
		}
		x := 1 + K*cnt
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 {
				m := (strength[i] + x - 1) / x
				next := state | 1<<i
				dp[next] = min(dp[next], dp[state]+m)
			}
		}
		
	}

	return dp[N-1]
}
