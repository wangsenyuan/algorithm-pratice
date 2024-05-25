package p526

func countArrangement1(N int) int {
	nums := make([]int, N)

	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}

	var check func(n int) int

	check = func(n int) int {
		if n == 0 {
			return 1
		}
		res := 0

		for i := 0; i < n; i++ {
			if nums[i]%n == 0 || n%nums[i] == 0 {
				nums[i], nums[n-1] = nums[n-1], nums[i]
				res += check(n - 1)
				nums[i], nums[n-1] = nums[n-1], nums[i]
			}
		}
		return res
	}

	return check(N)
}

func countArrangement(n int) int {
	N := 1 << n

	dp := make([]int, N)

	dp[0] = 1

	for state := 0; state < N; state++ {
		var sz int
		for i := 0; i < n; i++ {
			sz += (state >> i) & 1
		}
		sz++
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 && ((i+1)%sz == 0 || sz%(i+1) == 0) {
				dp[state|(1<<i)] += dp[state]
			}
		}
	}

	return dp[N-1]
}
