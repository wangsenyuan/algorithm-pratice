package p2294

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)

	var res int

	n := len(nums)

	for i := 0; i < n; {
		res++
		j := i

		for i < n && nums[i]-nums[j] <= k {
			i++
		}
	}

	return res
}
