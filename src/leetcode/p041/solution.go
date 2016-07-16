package main

import "fmt"

func main() {
	nums := []int{3, 4, -1, 1}
	r := firstMissingPositive(nums)
	fmt.Printf("%d is missing in %v\n", r, nums)
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}

	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 2
		}
	}

	for _, x := range nums {
		x = abs(x)
		if x <= n {
			nums[x-1] = -abs(nums[x-1])
		}
	}

	for i, x := range nums {
		if x > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
