package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	sort.Ints(nums)

	idx := make(map[int]int)

	var res int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			idx[nums[i]] = i
		}

		if j, found := idx[nums[i]-1]; found {
			if i-j+1 > res {
				res = i - j + 1
			}
		}
	}

	return res
}
