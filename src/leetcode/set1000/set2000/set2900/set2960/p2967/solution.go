package p2967

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	var res [][]int
	n := len(nums)
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		res = append(res, nums[i:i+3])
	}

	return res
}
