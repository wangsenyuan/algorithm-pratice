package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-4, -3}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	n := len(nums)
	res := math.MinInt32
	prod := 1

	updateRes := func(i int) {
		prod *= nums[i]
		if prod > res {
			res = prod
		}
		if nums[i] == 0 {
			prod = 1
		}
	}

	for i := 0; i < n; i++ {
		updateRes(i)
	}
	prod = 1
	for i := n - 1; i >= 0; i-- {
		updateRes(i)
	}
	return res
}
