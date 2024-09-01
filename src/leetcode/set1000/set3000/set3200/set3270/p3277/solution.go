package p3277

func maximumSubarrayXor(nums []int, queries [][]int) []int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			dp[l][r] = dp[l+1][r] ^ dp[l][r-1]
		}
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			dp[l][r] = max(dp[l+1][r], dp[l][r-1], dp[l][r])
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = dp[l][r]
	}
	return ans
}
