package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(search([]int{1, 3, 1, 1, 1}, 3))
}

func search(nums []int, target int) bool {
	p := - 1

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {
			p = i
			break
		}
	}

	index := -1
	if p < 0 {
		index = sort.SearchInts(nums, target)
	} else if nums[0] <= target {
		index = sort.SearchInts(nums[0:p+1], target)
	} else {
		index = p + 1 + sort.SearchInts(nums[p+1:], target)
	}

	if index < 0 || index >= len(nums) || nums[index] != target {
		return false
	}
	return true
}
