package p2357

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	cur := 0
	var res int
	n := len(nums)
	for i := 0; i < n; {
		for i < n && nums[i] == cur {
			i++
		}
		if i < n {
			res++
			cur = nums[i]
		}
	}

	return res
}
