package main

import (
	"fmt"
	"sort"
)

/**
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
Show Company Tags
Show Tags
Show Similar Problems
*
*/

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	for _, xs := range result {
		fmt.Printf("%v\n", xs)
	}
	threeSum([]int{0, 0, 0})
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	//fmt.Printf("%v\n", nums)
	result := make([][]int, 0, 10)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		c := nums[i]
		x := twoSum(nums, -c, i+1)
		result = append(result, x...)
	}

	return result
}

func twoSum(nums []int, c int, start int) [][]int {
	//fmt.Printf("try to find %d from %d\n", c, start)
	i, j := start, len(nums)-1
	result := make([][]int, 0, 10)
	for i < j {
		if i > start && nums[i] == nums[i-1] {
			i++
			continue
		}
		if j < len(nums)-1 && nums[j] == nums[j+1] {
			j--
			continue
		}
		sum := nums[i] + nums[j]
		//fmt.Printf("sum is %d\n", sum)
		if sum > c {
			j--
		} else if sum < c {
			i++
		} else {
			result = append(result, []int{-c, nums[i], nums[j]})
			i++
			j--
		}
	}
	//fmt.Printf("%d\n", len(result))
	return result
}
