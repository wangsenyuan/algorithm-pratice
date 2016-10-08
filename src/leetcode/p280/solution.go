package main

import "fmt"

func main() {
	nums := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 0 {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
}
