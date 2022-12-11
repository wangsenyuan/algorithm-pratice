package p2499

import "sort"

func longestSquareStreak(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	dp := make(map[int]int)
	var best int
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		if dp[x*x] > 0 {
			dp[x] = dp[x*x] + 1
		} else {
			dp[x] = 1
		}
		best = max(best, dp[x])
	}
	if best < 2 {
		return -1
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
