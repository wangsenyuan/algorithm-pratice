package p2576

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	// 2 * nums[i]<= nums[j]
	// 比如 1， 2， 3，4
	// 这是一个最大流的问题
	n := len(nums)
	mid := (n + 1) / 2
	j := mid
	var res int
	for i := 0; i < mid; i++ {
		for j < n && nums[i]*2 > nums[j] {
			j++
		}
		if j < n {
			res += 2
			j++
		}
	}
	return res
}
