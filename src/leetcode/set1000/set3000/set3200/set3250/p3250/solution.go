package p3250

import "slices"

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}
func sub(a, b int) int {
	return add(a, mod-b)
}

func countOfPairs(nums []int) int {
	mx := slices.Max(nums)

	dp := make([]int, mx+2)
	for i := 0; i <= nums[0]; i++ {
		dp[i] = 1
	}
	n := len(nums)
	diff := make([]int, mx+2)
	for i := 1; i < n; i++ {
		if nums[i-1] >= nums[i] {
			sum := dp[0]
			for x := 1; x <= nums[i]; x++ {
				sum = add(sum, dp[x])
				dp[x] = sum
			}
		} else {
			clear(diff)
			dist := nums[i] - nums[i-1]
			// x + u = nums[i-1]
			// y + v = nums[i]
			// x <= y
			// u >= v => nums[i-1] - x >= nums[i] - y
			// x <= y - dist
			for x := 0; x <= nums[i-1]; x++ {
				diff[x+dist] = dp[x]
			}

			for y := 1; y <= nums[i]; y++ {
				diff[y] = add(diff[y], diff[y-1])
			}

			copy(dp, diff)
		}

		clear(dp[nums[i]+1:])
	}

	var res int
	for i := 0; i <= mx; i++ {
		res = add(res, dp[i])
	}

	return res
}
