package p2770

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// dp[i][0] 表示最后一个来自nums1， dp[i][1] 表示最后一个来自nums2的最大长度
	dp := make([]int, 2)
	dp[0] = 1
	dp[1] = 1
	ans := 1
	for i := 1; i < n; i++ {
		a, b := 1, 1
		if nums1[i] >= nums1[i-1] {
			a = max(a, dp[0]+1)
		}
		if nums1[i] >= nums2[i-1] {
			a = max(a, dp[1]+1)
		}
		if nums2[i] >= nums1[i-1] {
			b = max(b, dp[0]+1)
		}
		if nums2[i] >= nums2[i-1] {
			b = max(b, dp[1]+1)
		}
		dp[0] = a
		dp[1] = b
		ans = max(ans, max(dp[0], dp[1]))
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
