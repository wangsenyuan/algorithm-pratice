package p2740

import "sort"

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	res := nums[n-1] - nums[0]

	for i := 0; i+1 < n; i++ {
		res = min(res, nums[i+1]-nums[i])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
