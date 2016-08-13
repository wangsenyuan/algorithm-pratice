package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &nums[j])
		}

		if play(m, nums) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func play(m int, nums []int) bool {
	sort.Ints(nums)
	n := len(nums)

	var doPlay func(i int, left int) bool

	doPlay = func(i int, left int) bool {
		for j := i; j < n; j++ {
			if nums[j] == left {
				return true
			} else if nums[j] < left {
				return doPlay(j+1, left-nums[j]) || doPlay(j+1, left)
			}
		}

		return false
	}

	return doPlay(0, m)
}
