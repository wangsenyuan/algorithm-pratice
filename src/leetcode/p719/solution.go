package main

import (
	"sort"
	"fmt"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)

	maxDist := nums[n-1] - nums[0]

	count := func(dist int) int {
		cnt := 0
		for i := 0; i < n; i++ {
			j := sort.Search(n, func(j int) bool {
				return nums[j] > nums[i]+dist
			})
			cnt += j - i - 1
		}

		return cnt
	}

	res := sort.Search(maxDist, func(dist int) bool {
		return count(dist) >= k
	})

	return res
}

func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 1))
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
}
