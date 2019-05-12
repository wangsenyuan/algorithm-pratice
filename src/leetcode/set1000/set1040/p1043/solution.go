package p1043


func maxSumAfterPartitioning(A []int, K int) int {

	// dp[i] = max sum ending at i
	n := len(A)

	dp := make([]int, n + 1)

	for i := 0; i < n; i++ {
		//dp[i+1]
		dp[i+1] = dp[i] + A[i]
		num := A[i]
		for k := 1; k < K && i - k >= 0; k++ {
			j := i - k
			num = max(num, A[j])
			dp[i+1] = max(dp[i+1], dp[j] + num * (k + 1))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}