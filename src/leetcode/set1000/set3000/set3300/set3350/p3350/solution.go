package p3350

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	best := 1
	fp := 1
	for i := n - 2; i > 0; i-- {
		best = max(best, min(dp[i], fp))
		if nums[i] < nums[i+1] {
			fp++
		} else {
			fp = 1
		}
	}
	return best
}
