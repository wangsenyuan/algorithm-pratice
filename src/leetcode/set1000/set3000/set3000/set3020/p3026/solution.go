package p3026

import "sort"

func triangleType(nums []int) string {
	sort.Ints(nums)

	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}

	if nums[0] == nums[2] {
		return "equilateral"
	}

	if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}

	return "scalene"
}
