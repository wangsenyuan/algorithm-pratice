package p2341

import "sort"

func numberOfPairs(nums []int) []int {
	sort.Ints(nums)

	var res int

	for i, j := 1, 0; i <= len(nums); i++ {
		if i == len(nums) || nums[i] > nums[i-1] {
			res += (i - j) / 2
			j = i
		}
	}

	return []int{res, len(nums) - 2*res}
}
