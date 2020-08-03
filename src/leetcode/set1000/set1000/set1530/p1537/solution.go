package p1537

const MOD = 1000000007

func maxSum(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([]int64, m+1)
	fp := make([]int64, n + 1)

	var i, j int

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			dp[i+1] = dp[i] + int64(nums1[i])
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			fp[j+1] = fp[j] + int64(nums2[j])
			j++
			continue
		}

		// ===
		x := max(dp[i], fp[j])
		dp[i+1] = x + int64(nums1[i])
		fp[j+1] = x + int64(nums2[j])
		i++
		j++
	}

	for i < m {
		dp[i+1] = dp[i] + int64(nums1[i])
		i++
	}
	for j < n {
		fp[j+1] = fp[j] + int64(nums2[j])
		j++
	}

	x := max(dp[m], fp[n])

	return int(x % MOD)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}