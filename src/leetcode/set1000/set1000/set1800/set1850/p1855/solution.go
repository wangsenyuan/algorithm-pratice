package p1855

func maxDistance(nums1 []int, nums2 []int) int {
	// nums[i] >= nums[i-1]
	// so if some j, that nums1[i] <= nums2[j], then >= i, all <= j
	i := len(nums1) - 1
	j := len(nums2) - 1
	var best int
	for j >= 0 {
		for i > 0 && nums1[i-1] <= nums2[j] {
			i--
		}

		if nums1[i] <= nums2[j] {
			best = max(best, j-i)
		}

		j--
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
