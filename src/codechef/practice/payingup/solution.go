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
	sums := make([]int, n+1)
	for i := range nums {
		sums[i+1] = sums[i] + nums[i]
	}

	var doPlay func(i int, left int) bool

	doPlay = func(i int, left int) bool {
		if left == 0 || left == sums[n]-sums[i] {
			return true
		}
		if left < nums[i] {
			return false
		}
		if left > sums[n]-sums[i] {
			return false
		}

		return doPlay(i+1, left-nums[i]) || doPlay(i+1, left)
	}

	return m > 0 && doPlay(0, m)
}
