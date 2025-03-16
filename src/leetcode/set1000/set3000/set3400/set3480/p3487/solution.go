package p3487

import "slices"

func maxSum(nums []int) int {
	pos := make(map[int]int)
	var sum int
	for _, num := range nums {
		if num >= 0 {
			pos[num]++
			if pos[num] == 1 {
				sum += num
			}
		}
	}
	if len(pos) > 0 {
		return sum
	}
	return slices.Max(nums)
}
