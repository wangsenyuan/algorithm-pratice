package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {

	var find func(nums []int, S int) int

	find = func(nums []int, S int) int {
		if len(nums) == 0 {
			if S == 0 {
				return 1
			}
			return 0
		}

		return find(nums[1:], S-nums[0]) + find(nums[1:], S+nums[0])
	}

	return find(nums, S)
}
