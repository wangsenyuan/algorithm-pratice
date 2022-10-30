package p2452

import "sort"

func destroyTargets(nums []int, space int) int {
	sort.Ints(nums)
	// space is the different gap
	// c could be >= 0
	dp := make(map[int]int)
	// cnt := make(map[int]int)

	n := len(nums)
	var x, ans int
	for i := n - 1; i >= 0; i-- {
		r := nums[i] % space
		dp[r]++
		if dp[r] >= x {
			ans = nums[i]
			x = dp[r]
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
