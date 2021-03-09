package p1787

func minChanges(nums []int, k int) int {
	n := len(nums)

	// xor[i:i+k] = 0
	// xor[i+1:i+k] = nums[i]
	// xor[i+1:i+1+k] = nums[i] ^ nums[i+k] = 0
	// nums[i] == nums[i+k]
	var res int
	if k == 1 {
		for i := 0; i < n; i++ {
			if nums[i] != 0 {
				res++
			}
		}
		return res
	}
	size := make([]int, k)
	for i := 0; i < n; i++ {
		size[i%k]++
	}
	// 前k个元素决定了结果
	// dp[N] = 使得前i个元素xor为state时的更改数,保留数
	N := 1 << 10
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = n
		}
	}

	count := func(p int) map[int]int {
		mem := make(map[int]int)
		for i := p; i < n; i += k {
			mem[nums[i]]++
		}
		return mem
	}

	dp[0][0] = 0
	var p int
	for i := 0; i < k; i++ {
		q := p ^ 1
		var best = n
		for cur := 0; cur < N; cur++ {
			best = min(best, dp[p][cur])
			dp[q][cur] = n
		}

		mem := count(i)

		for cur := 0; cur < N; cur++ {
			// we can use some number x, that x xor y (which get best before) to get cur value
			dp[q][cur] = min(dp[q][cur], size[i]+best)

			for num, v := range mem {
				// if we take num
				dp[q][cur^num] = min(dp[q][cur^num], size[i]-v+dp[p][cur])
			}
		}

		p = q
	}
	return dp[p][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
