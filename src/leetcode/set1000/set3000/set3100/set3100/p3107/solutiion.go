package p3107

import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)

	n := len(nums)

	median := n / 2
	// n = 2, median = 1

	if nums[median] == k {
		return 0
	}

	if nums[median] > k {
		// 减少前半部分的数
		res := int64(nums[median] - k)
		nums[median] = k
		for i := median - 1; i >= 0; i-- {
			diff := max(nums[i]-nums[i+1], 0)
			res += int64(diff)
			nums[i] -= diff
		}
		return res
	}
	// 增加后半部分的值
	res := int64(k - nums[median])
	nums[median] = k
	for i := median + 1; i < n; i++ {
		diff := max(nums[i-1]-nums[i], 0)
		res += int64(diff)
		nums[i] += diff
	}

	return res
}
