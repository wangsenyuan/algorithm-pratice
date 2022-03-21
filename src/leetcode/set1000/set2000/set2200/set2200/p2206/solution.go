package p2206

import "sort"

func divideArray(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}
