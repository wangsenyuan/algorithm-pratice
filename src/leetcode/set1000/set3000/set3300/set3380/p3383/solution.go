package p3383

import "math"

const inf = 1 << 60

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	// 前缀和

	pref := make([]int, n+1)
	for i, x := range nums {
		pref[i+1] = pref[i] + x
	}
	res := math.MinInt64

	// 如果k很小，怎么办？如果k很大怎么办？
	// 感觉可以用dp
	dp := make([]int, n+1)
	dp[0] = 0
	for i := k; i <= n; i++ {
		dp[i] = pref[i] - pref[i-k] + max(0, dp[i-k])
		res = max(res, dp[i])
	}

	return int64(res)
}
