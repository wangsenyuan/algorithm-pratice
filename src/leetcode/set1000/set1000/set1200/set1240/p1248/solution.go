package p1248

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] & 1
	}

	var sum int
	var res int
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 0; i < n; i++ {
		sum += nums[i]

		if nums[i] == 1 {
			dp[sum] = i + 1
		}

		if sum >= k {
			res += dp[sum-k+1] - dp[sum-k]
		}

	}

	return res
}
