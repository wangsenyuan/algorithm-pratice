package p1985

import "sort"

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := nums[i]
		b := nums[j]
		if len(a) > len(b) {
			return true
		}
		if len(a) == len(b) && a > b {
			return true
		}
		return false
	})
	return nums[k-1]
}
