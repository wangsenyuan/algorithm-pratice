package p1615

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n; i++ {
		cur := n - i
		if nums[i] >= cur {
			if i == 0 || nums[i-1] < cur {
				return cur
			}
		}
	}
	return -1
}
