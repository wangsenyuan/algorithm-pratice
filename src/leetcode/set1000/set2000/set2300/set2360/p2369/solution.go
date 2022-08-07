package p2369

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	//  means a valid partition ending at i
	for i := 2; i <= n; i++ {
		// check wether a valid partitiion can end at i
		if nums[i-2] == nums[i-1] {
			dp[i] = dp[i-2]
		}
		if i >= 3 {
			if nums[i-3] == nums[i-2] && nums[i-2] == nums[i-1] {
				dp[i] = dp[i] || dp[i-3]
			}
			if nums[i-2]-nums[i-3] == 1 && nums[i-1]-nums[i-2] == 1 {
				dp[i] = dp[i] || dp[i-3]
			}
		}
	}

	return dp[n]
}
