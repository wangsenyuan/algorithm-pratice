package p2741

const MOD = 1_000_000_007

func specialPerm(nums []int) int {
	n := len(nums)
	// n <= 14

	// dp[mask][i] 表示 排列完state表示的数后， 最后一位是i时的计数
	T := 1 << n

	dp := make([][]int, T)
	for i := 0; i < T; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for state := 1; state < T; state++ {
		for i := 0; i < n; i++ {
			if dp[state][i] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if state&(1<<j) > 0 || nums[i]%nums[j] != 0 && nums[j]%nums[i] != 0 {
					continue
				}
				dp[state|(1<<j)][j] = add(dp[state|(1<<j)][j], dp[state][i])
			}
		}
	}
	var res int
	for i := 0; i < n; i++ {
		res = add(res, dp[T-1][i])
	}

	return res
}

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
