package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closet := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := len(nums) - 1
		for k := i + 1; k < j; {
			sum := nums[i] + nums[k] + nums[j]
			if sum == target {
				return target
			}
			if sum > target {
				j--
			} else if sum < target {
				k++
			}
			diff := abs(sum - target)
			if closet == math.MaxInt32 || diff < abs(closet-target) {
				closet = sum
			}
		}
	}

	return closet
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
