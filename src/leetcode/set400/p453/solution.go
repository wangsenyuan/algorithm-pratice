package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}

func minMoves(nums []int) int {
	sum := 0
	min := math.MaxInt32

	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}

	return sum - min*len(nums)
}
