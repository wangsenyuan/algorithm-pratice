package p3132

import "sort"

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	m := len(nums2)
	n := len(nums1)
	// nums2[0] = nums1[0] + x
	// 或者 nums2[0] = nums1[1] + x
	// 或者nums2[0] = nums1[2] + x

	check := func(x int) bool {
		for i, j := 0, 0; i < m; i++ {
			for j < n && nums2[i]-nums1[j] != x {
				j++
				continue
			}
			if j == n {
				return false
			}
			j++
		}
		return true
	}

	res := 1 << 60
	for i := 0; i <= 2; i++ {
		x := nums2[0] - nums1[i]
		if check(x) {
			res = min(res, x)
		}
	}

	return res
}
