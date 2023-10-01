package p2875

import "sort"

func minOperations(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var res int
	for i := 0; i < n; {
		j := i

		for i < n && nums[i] == nums[j] {
			i++
		}
		if i-j == 1 {
			return -1
		}
		x, y := (i-j)/3, (i-j)%3
		if y == 0 {
			res += x
		} else {
			res += x + 1
		}
	}
	return res
}
