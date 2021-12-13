package p2104

import "math"

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	var res int64
	for i := 0; i < n; i++ {
		a := math.MaxInt32
		b := math.MinInt32
		for j := i; j < n; j++ {
			if nums[j] > b {
				b = nums[j]
			}
			if nums[j] < a {
				a = nums[j]
			}
			res += int64(b) - int64(a)
		}
	}

	return res
}
