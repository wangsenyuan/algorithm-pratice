package p2784

import "sort"

func isGood(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i-1] != i {
			return false
		}
	}
	if nums[n-1] != n-1 {
		return false
	}
	return true
}
