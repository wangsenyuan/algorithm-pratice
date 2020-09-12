package c

func minimumOperations(leaves string) int {
	n := len(leaves)

	var res int

	if leaves[0] == 'y' {
		res++
	}

	cnt := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		cnt[i] = cnt[i+1]
		if leaves[i] == 'r' {
			cnt[i]++
		}
	}

	dp := make([][]int, n)
	dp[0] = make([]int, 2)

	var best = n

	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2)
		if leaves[i] == 'y' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		}
		if i < n-1 {
			best = min(best, dp[i][1]+(n-(i+1)-cnt[i+1]))
		}
	}

	return best + res
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
