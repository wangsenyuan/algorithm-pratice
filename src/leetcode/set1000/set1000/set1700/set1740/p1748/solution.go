package p1748

import "sort"

func sumOfUnique(nums []int) int {
	var sum int

	sort.Ints(nums)

	for i := 0; i < len(nums); {
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
		if i == j+1 {
			sum += nums[j]
		}
	}

	return sum
}
