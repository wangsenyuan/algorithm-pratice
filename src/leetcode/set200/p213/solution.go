package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 1}))
}

func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := robHourse(nums, 0)
	b := robHourse(nums, 1)
	return max(a, b)
}

func robHourse(nums []int, from int) int {
	a := 0
	b := 0

	for i := 0; i < len(nums)-1; i++ {
		j := (from + i) % len(nums)

		c := a
		if b > c {
			c = b
		}
		b = a + nums[j]
		a = c
	}
	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
