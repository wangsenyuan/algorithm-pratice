package p1967

import "sort"

func rearrangeArray(nums []int) []int {
	// let a, b, c
	// 2 * b = a + c
	// swap a, b => b, a, c, because a , b, c unique
	sort.Ints(nums)
	n := len(nums)

	res := make([]int, 0, n)
	m := (n + 1) / 2

	for i := 0; i < m; i++ {
		res = append(res, nums[i])
		if i+m < n {
			res = append(res, nums[i+m])
		}
	}

	return res
}
