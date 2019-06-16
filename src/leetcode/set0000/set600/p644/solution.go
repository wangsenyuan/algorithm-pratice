package main

import (
	"math"
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	min, max := math.MaxFloat64, float64(math.MinInt32)

	for _, x := range nums {
		num := float64(x)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	prevMid := max
	error := 1.0

	for error > 0.00001 {
		mid := float64(min + 0.5*float64(max-min))
		if check(nums, mid, k) {
			min = mid
		} else {
			max = mid
		}

		error = math.Abs(prevMid - mid)
		prevMid = mid
	}

	return min
}

func check(nums []int, val float64, k int) bool {
	sum := 0.0

	for i := 0; i < k; i++ {
		sum += float64(nums[i]) - val
	}

	if sum >= 0 {
		return true
	}

	prev, min := 0.0, 0.0

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - val
		prev += float64(nums[i-k]) - val
		min = math.Min(prev, min)
		if sum >= min {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
