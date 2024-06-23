package p3194

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	var res = math.MaxFloat64

	sort.Ints(nums)

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		avg := float64(nums[i]+nums[j]) / 2
		res = math.Min(res, avg)
	}

	return res
}
