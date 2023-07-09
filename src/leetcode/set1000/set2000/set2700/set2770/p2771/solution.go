package p2771

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	// dp[i] = 到达i处的最大值
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if abs(nums[i]-nums[j]) <= target && dp[j] >= 0 {
				if dp[i] < 0 || dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n-1]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
