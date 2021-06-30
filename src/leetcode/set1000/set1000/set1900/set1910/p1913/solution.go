package p1913

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	a, b := nums[0], nums[1]
	c, d := nums[n-2], nums[n-1]

	return c*d - a*b
}
