package p3312

import (
	"slices"
	"sort"
)

func gcdValues(nums []int, queries []int64) []int {
	a := slices.Max(nums)
	freq := make([]int, a+1)
	for _, num := range nums {
		freq[num]++
	}

	// dp[i]表示有多少个对以i为gcd
	dp := make([]int, a+1)
	for i := a; i > 0; i-- {
		var cnt int
		for j := i; j <= a; j += i {
			cnt += freq[j]
			dp[i] -= dp[j]
		}
		dp[i] += cnt * (cnt - 1) / 2
	}

	for i := 1; i <= a; i++ {
		dp[i] += dp[i-1]
	}

	ans := make([]int, len(queries))

	for i, k := range queries {
		j := sort.Search(a+1, func(j int) bool {
			return dp[j] > int(k)
		})
		ans[i] = j
	}

	return ans
}
