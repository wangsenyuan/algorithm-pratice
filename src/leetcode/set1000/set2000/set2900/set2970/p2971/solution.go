package p2971

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	var res int64
	sum := nums[0] + nums[1]
	for i := 2; i < len(nums); i++ {
		if sum > nums[i] {
			res = int64(sum) + int64(nums[i])
		}
		sum += nums[i]
	}

	if res == 0 {
		return -1
	}

	return res
}
