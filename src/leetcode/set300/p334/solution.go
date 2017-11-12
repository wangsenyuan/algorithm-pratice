package main

import "fmt"

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
}

func increasingTriplet(nums []int) bool {
	i, j := -1, -1

	for k, num := range nums {
		if i == -1 || num <= nums[i] {
			i = k
		} else if j == -1 || num <= nums[j] {
			j = k
		} else {
			return true
		}
	}

	return false
}
