package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 0, 1, 3}
	target := 2
	fmt.Println(threeSumSmaller(nums, target))
}

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)

	res := 0
	for i := 0; i < len(nums)-2; i++ {
		res += twoSumSmaller(nums[i+1:], target-nums[i])
	}

	return res
}

func twoSumSmaller(nums []int, target int) int {
	i, j := 0, len(nums)-1

	res := 0
	for i < j {
		if nums[i]+nums[j] < target {
			res += j - i
			i++
			continue
		}

		j--
	}

	return res
}
