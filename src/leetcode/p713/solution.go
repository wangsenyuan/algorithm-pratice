package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	product := 1

	ans := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		for j < len(nums) && product*nums[j] < k {
			product *= nums[j]
			j++
		}

		if j >= i {
			ans += j - i
		}
		product /= nums[i]
	}

	return ans
}

func main() {
	//fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 80))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}
