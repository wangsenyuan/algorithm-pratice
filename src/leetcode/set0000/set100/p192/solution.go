package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	a, b := 0, 0

	for _, num := range nums {
		c := a
		if b > a {
			c = b
		}
		b = a + num
		a = c
	}

	if a > b {
		return a
	}
	return b
}
