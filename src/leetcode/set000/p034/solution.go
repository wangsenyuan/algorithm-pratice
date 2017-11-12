package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	i := sort.SearchInts(nums, target)
	if i < 0 || i >= len(nums) || nums[i] != target {
		return []int{-1, -1}
	}

	j := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})

	return []int{i, j - 1}
}

func main() {
	r := searchRange([]int{1}, 0)
	fmt.Printf("%d %d\n", r[0], r[1])
}
