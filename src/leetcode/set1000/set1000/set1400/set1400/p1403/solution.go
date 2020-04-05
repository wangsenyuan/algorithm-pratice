package p1403

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	reverse(nums)

	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	var cur int

	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if cur > sum-cur {
			return nums[:i+1]
		}
	}

	return nums
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
