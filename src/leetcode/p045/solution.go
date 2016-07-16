package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	r := jump(nums)
	fmt.Printf("%d steps for %v\n", r, nums)
}

func jump(nums []int) int {
	jp := 0
	last, curr := 0, 0

	for i, x := range nums {
		if i > last {
			last = curr
			jp++
		}
		if i+x > curr {
			curr = i + x
		}
	}

	return jp
}
