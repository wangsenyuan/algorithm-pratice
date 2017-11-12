package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	//j will stop before or at i + 1 anyway
	//swap i & j
	if i >= 0 {
		// nums[i] is the first number from right, nums[i] < nums[i+1]
		j := len(nums)
		for j > i+1 && nums[j-1] <= nums[i] {
			j--
		}
		nums[i], nums[j-1] = nums[j-1], nums[i]
	}

	//swap nums from i + 1 to right most
	n := len(nums)
	for k := 1; i+k < n-k; k++ {
		nums[i+k], nums[n-k] = nums[n-k], nums[i+k]
	}
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
