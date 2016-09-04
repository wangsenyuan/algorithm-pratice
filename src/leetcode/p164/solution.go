package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 5, 2, 9}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	a, b := math.MaxInt32, math.MinInt32

	for _, num := range nums {
		if num < a {
			a = num
		}
		if num > b {
			b = num
		}
	}
	n := len(nums)

	avg := (b-a)/(n-1) + 1

	bucks := make([]int, 2*n)

	for i := range bucks {
		bucks[i] = -1
	}

	for _, num := range nums {
		i := (num - a) / avg
		if bucks[2*i] == -1 || bucks[2*i] > num {
			bucks[2*i] = num
		}

		if bucks[2*i+1] == -1 || bucks[2*i+1] < num {
			bucks[2*i+1] = num
		}
	}
	start := bucks[1]
	gap := bucks[1] - bucks[0]

	for i := 1; i < n; i++ {
		if bucks[2*i] == -1 {
			continue
		}

		if bucks[2*i]-start > gap {
			gap = bucks[2*i] - start
		}
		start = bucks[2*i+1]
	}

	return gap
}
