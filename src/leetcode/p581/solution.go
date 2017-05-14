package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}

func findUnsortedSubarray(nums []int) int {
	std := make([]int, len(nums))
	copy(std, nums)
	sort.Ints(std)

	var i = 0
	for i < len(nums) && std[i] == nums[i] {
		i++
	}

	var j = len(nums) - 1
	for j > i && std[j] == nums[j] {
		j--
	}

	return j - i + 1
}
