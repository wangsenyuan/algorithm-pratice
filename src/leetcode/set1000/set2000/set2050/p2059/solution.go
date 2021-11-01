package p2059

func minimumOperations(nums []int, start int, goal int) int {
	n := len(nums)
	// n <= 1000
	dp := make([]int, 1001)
	for i := 0; i <= 1000; i++ {
		dp[i] = -1
	}
	dp[start] = 0
	fp := make([]int, 1001)
	gp := make([]int, 1001)
	for i := 0; i < n; i++ {
		for j := 0; j <= 1000; j++ {
			fp[j] = -1
			gp[j] = -1
		}
		num := nums[i]

		if num > 0 {
			// try adding num from j
			for j := num; j <= 1000; j++ {
				// j - num >= 0 && j - num <= 1000
				if dp[j-num] >= 0 {
					update(&fp[j], dp[j-num]+1)
				}
				if fp[j-num] >= 0 {
					update(&fp[j], fp[j-num]+1)
				}
			}
			// try sub num from j
			for j := 1000 - num; j >= 0; j-- {
				if dp[j+num] >= 0 {
					update(&gp[j], dp[j+num]+1)
				}
				if gp[j+num] >= 0 {
					update(&gp[j], gp[j+num]+1)
				}
			}
		} else if num < 0 {
			// adding num will make j decrease,
			for j := 1000 + num; j >= 0; j-- {
				if dp[j-num] >= 0 {
					update(&fp[j], dp[j-num]+1)
				}
				if fp[j-num] >= 0 {
					update(&fp[j], fp[j-num]+1)
				}
			}
			for j := -num; j <= 1000; j++ {
				if dp[j+num] >= 0 {
					update(&gp[j], dp[j+num]+1)
				}
				if gp[j+num] >= 0 {
					update(&gp[j], gp[j+num]+1)
				}
			}
		}

		for j := 0; j <= 1000; j++ {
			if dp[j] < 0 {
				continue
			}
			tmp := j ^ num
			if tmp >= 0 && tmp <= 1000 {
				update(&fp[tmp], dp[j]+1)
			}
		}

		for j := 0; j <= 1000; j++ {
			if fp[j] >= 0 {
				update(&dp[j], fp[j])
			}
			if gp[j] >= 0 {
				update(&dp[j], gp[j])
			}
		}
	}
	if goal <= 1000 && goal >= 0 {
		return dp[goal]
	}
	res := -1

	for i := 0; i < n; i++ {
		num := nums[i]
		for j := 0; j <= 1000; j++ {
			if dp[j] < 0 {
				continue
			}
			if num^j == goal {
				update(&res, dp[j]+1)
			}
			if j+num == goal {
				update(&res, dp[j]+1)
			}
			if j-num == goal {
				update(&res, dp[j]+1)
			}
		}
	}

	return res
}

func update(a *int, b int) {
	if *a < 0 || *a > b {
		*a = b
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
