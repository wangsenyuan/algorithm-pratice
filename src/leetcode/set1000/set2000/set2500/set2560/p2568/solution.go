package p2568

import "sort"

func minImpossibleOR(nums []int) int {
	sort.Ints(nums)

	var or int
	// 1, 2, 4, 8

	for _, num := range nums {
		if or+1 < num {
			return or + 1
		}
		or |= num
	}

	return or + 1
}
