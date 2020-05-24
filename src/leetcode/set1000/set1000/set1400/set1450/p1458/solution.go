package p1458

func maxDotProduct(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	// l := min(m, n)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = nums1[i] * nums2[j]
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+nums1[i]*nums2[j])
			}
		}
	}

	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
