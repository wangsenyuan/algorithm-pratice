package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 2, 5, 3, 5}
	fmt.Println(thirdMax(nums))
}
func thirdMax(nums []int) int {
	a, b, c := math.MinInt32, math.MinInt32, math.MinInt32

	for _, x := range nums {
		if x == a || x == b || x == c {
			continue
		}
		if x > a {
			a, b, c = x, a, b
		} else if x > b {
			b, c = x, b
		} else if x > c {
			c = x
		}
	}
	if len(nums) >= 3 {
		return c
	}

	return a
}
