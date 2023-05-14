package p2679

import "sort"

func matrixSum(nums [][]int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
	}

	m := len(nums[0])
	var res int
	for j := 0; j < m; j++ {
		cur := nums[0][j]
		for i := 1; i < n; i++ {
			if nums[i][j] > cur {
				cur = nums[i][j]
			}
		}
		res += cur
	}
	return res
}
