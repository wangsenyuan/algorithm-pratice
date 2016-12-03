package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	cnts := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		list := make([]int, 0)
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] != 0 {
				continue
			}
			if len(cnts[j]) > len(list) {
				list = cnts[j]
			}
		}
		cp := make([]int, len(list))
		copy(cp, list)

		cnts[i] = append(cp, nums[i])
	}

	res := make([]int, 0)

	for _, cnt := range cnts {
		if len(cnt) > len(res) {
			res = cnt
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n", largestDivisibleSubset([]int{1, 2, 4, 8, 72}))
}
