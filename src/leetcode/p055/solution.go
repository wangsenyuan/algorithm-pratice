package main

import "fmt"

func main() {
	fmt.Printf("can jump ? %v\n", canJump([]int{2, 3, 1, 1, 4}))
	fmt.Printf("can jump ? %v\n", canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	last := 0
	for i, num := range nums {
		if i > last {
			return false
		}
		if i+num > last {
			last = i + num
		}
	}

	return true
}
