package p1887

import "sort"

func reductionOperations(nums []int) int {
	sort.Ints(nums)

	var cnt int

	var res int

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt++
		}
		res += cnt
	}

	return res
}
