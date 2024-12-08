package p3375

import "sort"

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	var res int
	n := len(nums)
	if nums[n-1] < k {
		return -1
	}
	for i := n - 1; i >= 0; {
		j := i

		for i >= 0 && nums[i] == nums[j] {
			i--
		}
		if nums[j] == k {
			if i < 0 {
				return res
			}
			return -1
		}
		if nums[j] < k
		res++
	}
	return res
}
