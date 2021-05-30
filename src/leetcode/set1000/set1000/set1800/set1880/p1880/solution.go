package p1880

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)

	var best = 0

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i] + nums[j]
		if tmp > best {
			best = tmp
		}
	}
	return best
}
