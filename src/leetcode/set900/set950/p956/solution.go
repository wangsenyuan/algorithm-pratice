package p956

import (
	"math"
)

const NINF = math.MinInt32 / 2

func tallestBillboard(rods []int) int {
	n := len(rods)
	dp := make([]int, 5001)
	for i := 1; i < 5001; i++ {
		dp[i] = -10000
	}
	tmp := make([]int, 5001)
	for i := 0; i < n; i++ {
		x := rods[i]
		copy(tmp, dp)
		for d := 0; d+x < 5001; d++ {
			dp[d+x] = max(dp[d+x], tmp[d])
			dp[abs(d-x)] = max(dp[abs(d-x)], tmp[d]+min(d, x))
		}
	}
	return dp[0]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func tallestBillboard2(rods []int) int {
	n := len(rods)

	var dp func(i int, j int) int

	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, 10001)
		for j := 0; j < 10001; j++ {
			mem[i][j] = NINF
		}
	}

	dp = func(i int, j int) int {
		if i == n {
			if j == 5000 {
				return 0
			}
			return NINF
		}
		if mem[i][j] == NINF {
			ans := dp(i+1, j)
			ans = max(ans, dp(i+1, j-rods[i]))
			ans = max(ans, rods[i]+dp(i+1, j+rods[i]))
			mem[i][j] = ans
		}

		return mem[i][j]
	}

	return dp(0, 5000)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func tallestBillboard1(rods []int) int {
	n := len(rods)
	all := 1 << uint(n)
	sums := make([][]int, 5010)
	for i := 0; i < len(sums); i++ {
		sums[i] = make([]int, 0, 10)
	}
	var big int
	for mask := 1; mask < all; mask++ {
		// mask is the picked one
		var sum int

		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) > 0 {
				sum += rods[i]
			}
		}
		sums[sum] = append(sums[sum], mask)
		if sum > big {
			big = sum
		}
	}

	for x := big; x > 0; x-- {
		if len(sums) > 1 {
			for i := 0; i < len(sums[x])-1; i++ {
				mask1 := sums[x][i]
				for j := i + 1; j < len(sums[x]); j++ {
					mask2 := sums[x][j]
					if mask1&mask2 == 0 {
						return x
					}
				}
			}
		}
	}

	return 0
}
