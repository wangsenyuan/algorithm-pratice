package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1}
	rs := permuteUnique(nums)
	for _, r := range rs {
		fmt.Printf("%v\n", r)
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0, 10)
	rs = append(rs, copySlice(nums))
	last := false

	for !last {
		last = true
		for k := len(nums) - 1; k > 0; k-- {
			if nums[k-1] >= nums[k] {
				continue
			}
			last = false
			l := len(nums) - 1
			for nums[l] <= nums[k-1] {
				l--
			}
			nums[k-1], nums[l] = nums[l], nums[k-1]

			for a, b := k, len(nums)-1; a < b; a, b = a+1, b-1 {
				nums[a], nums[b] = nums[b], nums[a]
			}

			rs = append(rs, copySlice(nums))
		}
	}

	return rs
}

func copySlice(nums []int) []int {
	t := make([]int, len(nums))
	copy(t, nums)
	return t
}
