package p217

import "math"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	min, max := math.MaxInt32, math.MinInt32

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	if max-min+1 < len(nums) {
		return true
	}
	set := make([]bool, max-min+1)
	for _, num := range nums {
		i := num - min
		if set[i] {
			return true
		}
		set[i] = true
	}
	return false
}
