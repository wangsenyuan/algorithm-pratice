package p2787

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	dp := make([]int64, 2)
	for i := n - 1; i >= 0; i-- {
		cur := nums[i]
		if cur&1 == 0 {
			dp[0] = max(dp[0]+int64(cur), dp[1]-int64(x)+int64(cur))
		} else {
			dp[1] = max(dp[1]+int64(cur), dp[0]-int64(x)+int64(cur))
		}
	}
	return dp[nums[0]&1]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
