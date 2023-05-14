package p2680

func maximumOr(nums []int, k int) int64 {
	n := len(nums)

	dp := make([]int64, k+1)
	fp := make([]int64, k+1)
	var or int
	for i := 0; i < n; i++ {
		or |= nums[i]
		for j := 0; j <= k; j++ {
			fp[j] = 0
		}
		fp[0] = int64(or)
		tmp := int64(nums[i])

		for a := 0; a <= k; a++ {
			for b := 0; a+b <= k; b++ {
				fp[a+b] = max(fp[a+b], tmp|dp[b])
			}
			tmp *= 2
		}
		copy(dp, fp)
		for j := 1; j <= k; j++ {
			dp[j] = max(dp[j], dp[j-1])
		}
	}

	return dp[k]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
