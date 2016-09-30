package main

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	//right-most set bit, also is the right-most different set bit of a & b
	xor &= -xor
	a, b := 0, 0
	for _, num := range nums {
		if xor&num > 0 {
			a = num ^ a
		} else {
			b = num ^ b
		}
	}
	return []int{a, b}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
