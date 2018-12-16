package p960

func minDeletionSize(A []string) int {
	m := len(A)
	if m == 0 {
		return 0
	}
	n := len(A[0])
	if n == 0 {
		return 0
	}

	sorted := func(i, j int) bool {
		for k := 0; k < m; k++ {
			if A[k][i] > A[k][j] {
				return false
			}
		}
		return true
	}

	dp := make([]int, n)
	dp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if sorted(i, j) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	var best int

	for i := 0; i < n; i++ {
		best = max(best, dp[i])
	}
	return n - best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
