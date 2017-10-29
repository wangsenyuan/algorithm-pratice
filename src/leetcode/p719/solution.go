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
			j := i;
			for j < n && nums[j]-nums[i] <= dist {
				j++
			}
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
	//fmt.Println(smallestDistancePair([]int{1, 3, 1}, 1))
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
}
