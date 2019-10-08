package p1218

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)

	var best int

	for i := 0; i < len(arr); i++ {
		x := arr[i]

		if y, found := dp[x-difference]; found {
			dp[x] = max(dp[x], y+1)
		} else {
			dp[x] = max(dp[x], 1)
		}

		best = max(best, dp[x])
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
