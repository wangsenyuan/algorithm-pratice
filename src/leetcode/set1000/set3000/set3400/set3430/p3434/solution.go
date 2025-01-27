package p3434

import "slices"

func maxFrequency(nums []int, k int) int {
	// dp[i][x] = pref[i][k] - pref[i][x] 到i为止的最大值
	n := len(nums)
	x := 50
	dp := make([][]int, n+1)
	dp[0] = make([]int, x+1)
	pref := make([][]int, n+1)
	pref[0] = make([]int, x+1)

	// pref[i][k] - pref[i][num]
	best := make([]int, x+1)

	for i := range n {
		pref[i+1] = make([]int, x+1)
		copy(pref[i+1], pref[i])
		num := nums[i]
		pref[i+1][num]++
		dp[i+1] = make([]int, x+1)
		copy(dp[i+1], dp[i])
		dp[i+1][num] = max(dp[i+1][num], pref[i+1][num]+best[num])
		if num == k {
			for y := 1; y <= x; y++ {
				best[y] = max(best[y], pref[i+1][k]-pref[i+1][y])
			}
		} else {
			best[num] = max(best[num], pref[i+1][k]-pref[i+1][num])
		}
	}

	ans := slices.Max(dp[n])
	var cnt int
	for i := n - 1; i >= 0; i-- {
		if nums[i] == k {
			cnt++
		}
		for y := 1; y <= x; y++ {
			tmp := cnt + dp[i][y]
			ans = max(ans, tmp)
		}
	}

	return ans
}
