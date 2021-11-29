package p2088

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)

	var res []int

	for i, num := range nums {
		if num == target {
			res = append(res, i)
		}
	}
	return res
}
