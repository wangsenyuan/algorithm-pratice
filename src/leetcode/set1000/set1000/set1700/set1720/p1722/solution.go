package p1722

func minimumTimeRequired(jobs []int, k int) int {
	if k == 1 {
		return sum(jobs)
	}
	n := len(jobs)
	N := 1 << n
	// dp[i][state]
	X := make([]int, N)
	for state := 0; state < N; state++ {
		var tmp int
		for j := 0; j < n; j++ {
			if state&(1<<j) > 0 {
				tmp += jobs[j]
			}
		}
		X[state] = tmp
	}

	dp := make([]int, N)
	for i := 1; i < N; i++ {
		dp[i] = 1 << 30
	}

	for i := 0; i < k; i++ {
		for state := N - 1; state > 0; state-- {
			for cur := state; cur > 0; cur = (cur - 1) & state {
				dp[state] = min(dp[state], max(dp[state^cur], X[cur]))
			}
		}
	}
	return dp[N-1]
}

func sum(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
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
