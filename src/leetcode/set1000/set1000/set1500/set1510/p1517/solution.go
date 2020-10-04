package p1517

import "sort"

func numIdenticalPairs(nums []int) int {
	sort.Ints(nums)

	var res int

	var j int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] > nums[i-1] {
			cnt := i - j
			res += cnt * (cnt - 1) / 2
			j = i
		}
	}

	return res
}
