package p3049

import "sort"

func isPossibleToSplit(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums); {
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
		if i-j > 2 {
			return false
		}
	}
	return true
}
