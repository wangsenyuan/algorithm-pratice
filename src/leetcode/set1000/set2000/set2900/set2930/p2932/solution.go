package p2932

func minOperations(nums1 []int, nums2 []int) int {
	n := len(nums1)

	check := func(a, b int) int {
		// a is max for nums1, b is max for nums2

		var res int
		for i := n - 2; i >= 0; i-- {
			if nums1[i] <= a && nums2[i] <= b {
				// already good
				continue
			}
			// have to swap
			if nums2[i] <= a && nums1[i] <= b {
				res++
				continue
			}
			return -2
		}
		return res
	}

	x := check(nums1[n-1], nums2[n-1])
	y := check(nums2[n-1], nums1[n-1]) + 1

	if x >= 0 && y >= 0 {
		return min(x, y)
	}
	if y < 0 && x < 0 {
		return -1
	}
	if y < 0 {
		return x
	}
	return y
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
