package main

import "fmt"

func main() {
	nums := []int{3, 3, 1, 3}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	n := len(nums)

	i, j := 0, n-1

	for i < j && nums[i] == nums[j] {
		j--
	}

	for i < j {
		mid := (i + j) / 2
		if nums[mid] > nums[j] {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return nums[i]
}
