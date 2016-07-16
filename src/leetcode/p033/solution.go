package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1}
	i := search(nums, 0)
	fmt.Printf("%d\n", i)
}

func search(nums []int, target int) int {
	p := -1
	i, j := 0, len(nums) - 1
	for i < j && nums[i] > nums[j] {
		m := (i + j) / 2
		if i == m {
			p = i
			break
		}
		if nums[i] < nums[m] {
			i = m
			p = i
		} else {
			j = m
		}
	}
	index := -1
	if p < 0 {
		index = sort.SearchInts(nums, target)
	} else if nums[0] <= target {
		index = sort.SearchInts(nums[0:p + 1], target)
	} else {
		index = p + 1 + sort.SearchInts(nums[p + 1:], target)
	}

	if (index < 0 || index >= len(nums) || nums[index] != target) {
		return -1
	}
	return index
}
