package main

import "fmt"

func main() {
	nums := []int{1, 0}
	fmt.Printf("%d is missing from %v\n", missingNumber(nums), nums)
}

func missingNumber(nums []int) int {
	n := len(nums)

	for i := range nums {
		nums[i]++
	}

	for _, x := range nums {
		x = abs(x)
		if x <= n {
			nums[x-1] = -abs(nums[x-1])
		}
	}

	for i, x := range nums {
		if x > 0 {
			return i
		}
	}

	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
