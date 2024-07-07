package p3279

import "sort"

const H = 30

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)
	dp := make([][]int, n+1)
	dp[0] = make([]int, H)
	for j := 0; j < H; j++ {
		dp[0][j] = -1
	}
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, H)
		dp[i][0] = nums[i-1]
		for j := 1; j < H; j++ {
			if i >= 1<<(j-1) {
				dp[i][j] = dp[i][j-1] & dp[i-(1<<(j-1))][j-1]
			}
		}
	}

	getValue := func(l int, r int) int {
		ans := 1<<H - 1

		for i := H - 1; i >= 0; i-- {
			if (r-l)&(1<<i) > 0 {
				ans &= dp[r][i]
				r -= 1 << i
			}
		}

		return ans
	}

	var res int

	for i := 1; i <= n; i++ {
		a := sort.Search(i, func(j int) bool {
			return getValue(j, i) > k
		})
		b := sort.Search(i, func(j int) bool {
			return getValue(j, i) >= k
		})

		res += a - b
	}

	return int64(res)
}
