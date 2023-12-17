package p2965

import "sort"

func maxFrequencyScore(nums []int, k int64) int {
	sort.Ints(nums)
	n := len(nums)

	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(nums[i])
	}

	check := func(l int, r int) bool {
		i := (l + r) / 2
		a := int64(i-l)*int64(nums[i]) - (sum[i] - sum[l])
		b := (sum[r+1] - sum[i+1]) - int64(r-i)*int64(nums[i])
		return a+b <= k
		// find mid
	}

	var best int

	for l, r := 0, 0; r < n; r++ {
		for !check(l, r) {
			l++
		}
		best = max(best, r-l+1)
	}

	return best
}
