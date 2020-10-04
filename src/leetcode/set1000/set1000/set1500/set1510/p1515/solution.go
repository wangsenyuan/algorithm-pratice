package p1515

import "sort"

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	n := len(nums)
	sort.Ints(nums)
	var best int = nums[n-1] - nums[0]
	for x := 3; x >= 0; x-- {
		y := 3 - x
		a := nums[x]
		b := nums[n-1-y]
		if b-a < best {
			best = b - a
		}
	}

	return best
}
