package p2605

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	x := nums1[0]*10 + nums2[0]
	y := nums2[0]*10 + nums1[0]
	return min(x, y)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
