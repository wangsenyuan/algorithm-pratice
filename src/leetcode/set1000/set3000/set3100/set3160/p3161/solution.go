package p3161

import "sort"

func duplicateNumbersXOR(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); {
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
		if i-j == 2 {
			res ^= nums[j]
		}
	}
	return res
}
