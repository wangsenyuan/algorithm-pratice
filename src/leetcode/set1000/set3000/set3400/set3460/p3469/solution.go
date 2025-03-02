package p3469

import "slices"

const inf = 1 << 60

func minCost(nums []int) int {
	n := len(nums)
	if n < 3 {
		return slices.Max(nums)
	}
	dp := make([]int, n)

	for i := range n {
		dp[i] = inf
	}
	// n <= 1000
	// dp[i][j] 表示前缀处理完后，剩余的数时a[j]时的最优解
	dp[0] = max(nums[1], nums[2])
	dp[1] = max(nums[0], nums[2])
	dp[2] = max(nums[0], nums[1])
	for i := 3; i < n; i += 2 {
		// 每次要选这个数和前面的进行处置
		if i+1 == n {
			cur := inf
			// 只能和nums[i]比较, 也就是只剩两个数了
			for j := range i {
				cur = min(cur, dp[j]+max(nums[j], nums[i]))
			}
			return cur
		}
		// 还有两个数可以用, nums[i], nums[i+1]
		for j := range i {
			dp[i] = min(dp[i], dp[j]+max(nums[j], nums[i+1]))
			dp[i+1] = min(dp[i+1], dp[j]+max(nums[j], nums[i]))
			dp[j] += max(nums[i], nums[i+1])
		}
	}
	ans := inf
	for i := range n {
		ans = min(ans, dp[i]+nums[i])
	}

	return ans
}
