package p2673

func minIncrements(n int, cost []int) int {
	// n = (2 ** d) - 1
	h := (n + 1) / 2
	sum := make([]int, n+1)
	for i := h; i <= n; i++ {
		sum[i] = cost[i-1]
	}
	// sum[i] = the max path-sum from i to leaf j
	// dp[i] = min operations to make sub-tree rooted at i, to meet the requirements
	var res int
	for i := h - 1; i > 0; i-- {
		l := 2 * i
		r := l + 1
		diff := abs(sum[l] - sum[r])
		// add diff to the less-sum node
		res += diff
		sum[i] = max(sum[l], sum[r]) + cost[i-1]
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
