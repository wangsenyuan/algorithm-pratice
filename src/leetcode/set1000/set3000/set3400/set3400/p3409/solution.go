package p3409

import "slices"

func longestSubsequence(nums []int) int {
	// a[i] <= 300
	x1 := slices.Max(nums)

	// dp[i][x] = longest subsequence ending at i with difference x
	// dp[i][x] = max(dp[j][x] + 1) for all j < i and abs(nums[i] - nums[j]) =  x
	pos := make([]int, x1+1)
	for i := 0; i <= x1; i++ {
		pos[i] = -1
	}
	n := len(nums)
	dp := make([][]int, n)
	var ans int
	for i := 0; i < n; i++ {
		dp[i] = make([]int, x1)
		for d := 0; d < x1; d++ {
			if nums[i]-d > 0 && pos[nums[i]-d] >= 0 {
				j := pos[nums[i]-d]
				dp[i][d] = max(dp[i][d], dp[j][d]+1)
			}
			if nums[i]+d <= x1 && pos[nums[i]+d] >= 0 {
				j := pos[nums[i]+d]
				dp[i][d] = max(dp[i][d], dp[j][d]+1)
			}
		}
		for d := x1 - 2; d >= 0; d-- {
			dp[i][d] = max(dp[i][d], dp[i][d+1])
		}
		ans = max(ans, max(dp[i][0]))
		pos[nums[i]] = i
	}

	return ans + 1
}
