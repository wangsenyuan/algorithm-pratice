package p2592

import "sort"

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// [1,3,5,2,1,3,1]
	// [1,1,1,2,3,3,5]
	// [2,3,3,5,1,1,1]
	var res int

	for i, j := 0, 1; j < n; j++ {
		if nums[i] < nums[j] {
			res++
			i++
		}
	}

	return res
}
