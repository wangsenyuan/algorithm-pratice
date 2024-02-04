package p3025

import "math"

func maximumSubarraySum(nums []int, k int) int64 {

	mn := make(map[int]int64)

	var sum int64

	var best int64 = math.MinInt64

	for _, num := range nums {
		sum += int64(num)
		if v, ok := mn[num+k]; ok {
			best = max(best, sum-v)
		}
		if v, ok := mn[num-k]; ok {
			best = max(best, sum-v)
		}

		if v, ok := mn[num]; !ok || v > sum-int64(num) {
			mn[num] = sum - int64(num)
		}
	}

	if best == math.MinInt64 {
		return 0
	}

	return best
}
