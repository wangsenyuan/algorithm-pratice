package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	half := sum / 2
	sort.Ints(nums)

	var pick func(i int, sum int) bool

	pick = func(i int, sum int) bool {
		if sum == half {
			return true
		}
		if i < 0 {
			return false
		}

		if i*nums[i]+sum < half {
			return false
		}

		if sum+nums[i] <= half && pick(i-1, sum+nums[i]) {
			return true
		}
		return pick(i-1, sum)
	}

	return pick(len(nums)-1, 0)
}
