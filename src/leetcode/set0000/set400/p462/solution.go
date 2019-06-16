package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMoves2([]int{1, 0, 0, 8, 6}))
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		moves += nums[j] - nums[i]
	}

	return moves
}
