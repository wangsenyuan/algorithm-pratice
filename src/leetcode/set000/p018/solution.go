package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	result := fourSum(nums, 0)
	for _, x := range result {
		fmt.Printf("%v\n", x)
	}
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, 10)

	for i := len(nums) - 1; i > 2; i-- {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		xs := threeSum(nums[:i], target-nums[i])
		for _, x := range xs {
			result = append(result, append(x, nums[i]))
		}
	}

	return result
}

func threeSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)
	for i := len(nums) - 1; i > 1; i-- {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		xs := twoSum(nums[:i], target-nums[i])
		for _, x := range xs {
			result = append(result, append(x, nums[i]))
		}
	}

	return result
}

func twoSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)
	for i, j := 0, len(nums)-1; i < j; {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		if j < len(nums)-1 && nums[j] == nums[j+1] {
			j--
			continue
		}

		sum := nums[i] + nums[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			result = append(result, []int{nums[i], nums[j]})
			i++
			j--
		}
	}
	return result
}
