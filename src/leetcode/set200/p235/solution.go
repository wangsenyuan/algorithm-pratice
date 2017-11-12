package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	a := make([]int, n)
	a[0] = 1
	for i := 1; i < len(nums); i++ {
		a[i] = a[i-1] * nums[i-1]
	}

	p := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		a[i] = a[i] * p
		p = nums[i] * p
	}

	return a
}
