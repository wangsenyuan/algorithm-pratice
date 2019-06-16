package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	fmt.Println(findKthLargest(nums, 2))
}

func findKthLargest(nums []int, k int) int {
	left, right := make([]int, len(nums)-1), make([]int, len(nums)-1)
	var find func(nums []int, l int, i int) int

	find = func(nums []int, l int, i int) int {
		pivot := nums[0]
		a, b := 0, 0
		for j := 1; j < l; j++ {
			if nums[j] > pivot {
				left[a] = nums[j]
				a++
			} else {
				right[b] = nums[j]
				b++
			}
		}

		if a == i {
			return pivot
		}
		if a > i {
			return find(left, a, i)
		}

		return find(right, b, i-a-1)
	}

	return find(nums, len(nums), k-1)
}
