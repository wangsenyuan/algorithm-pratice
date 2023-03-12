package p2588

import "sort"

func maxScore(nums []int) int {
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	var sum int64
	var res int
	for _, num := range nums {
		sum += int64(num)
		if sum <= 0 {
			break
		}
		res++
	}

	return res
}
