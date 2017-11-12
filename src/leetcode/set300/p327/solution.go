package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 5, -1}
	fmt.Println(countRangeSum(nums, -2, 2))
}

func countRangeSum(nums []int, lower int, upper int) int {

	first := make([]int, len(nums)+1)
	for i := range nums {
		first[i+1] = first[i] + nums[i]
	}

	var count func(lo int, hi int) int

	count = func(lo, hi int) int {
		mid := (hi + lo) / 2
		if mid == lo {
			return 0
		}
		ans := count(lo, mid) + count(mid, hi)

		i, j := mid, mid
		for _, left := range first[lo:mid] {
			for i < hi && first[i]-left < lower {
				i++
			}

			for j < hi && first[j]-left <= upper {
				j++
			}
			ans += j - i
		}
		sort.Ints(first[lo:hi])
		return ans
	}

	return count(0, len(first))
}
