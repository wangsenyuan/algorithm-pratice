package p2569

import (
	"sort"
)

func minimizeSum(nums []int) int {
	n := len(nums)
	if n == 3 {
		return 0
	}
	sort.Ints(nums)
	a := nums[n-3] - nums[0]
	b := nums[n-2] - nums[1]
	c := nums[n-1] - nums[2]
	return min(a, min(b, c))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
