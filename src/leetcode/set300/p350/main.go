package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	nums3 := make([]int, 0, min(len(nums1), len(nums2)))

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); i++ {
		x := nums1[i]
		for j < len(nums2) && nums2[j] < x {
			j++
		}

		if j == len(nums2) {
			break
		}

		if nums2[j] > x {
			continue
		}

		nums3 = append(nums3, x)
		j++
	}

	return nums3
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
