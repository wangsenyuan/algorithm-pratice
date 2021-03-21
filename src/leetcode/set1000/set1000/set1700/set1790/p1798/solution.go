package p1798

func maxScore(nums []int) int {
	N := len(nums)
	//n := N >> 1
	// n <= 7
	mat := make([][]int, N)

	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
		for j := 0; j < N; j++ {
			mat[i][j] = gcd(nums[i], nums[j])
		}
	}
	dp := make([]int, 1<<N)
	for i := 0; i < 1<<N; i++ {
		dp[i] = -(1 << 30)
	}
	dp[0] = 0
	for mask := 3; mask < 1<<N; mask++ {
		// 2 ^^ 14
		cnt := digitCount(mask)
		if cnt%2 != 0 {
			continue
		}
		for i := 0; i < N; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			for j := i + 1; j < N; j++ {
				if mask&(1<<j) == 0 {
					continue
				}
				mask2 := mask ^ (1 << i) ^ (1 << j)
				dp[mask] = max(dp[mask], dp[mask2]+(cnt/2)*mat[i][j])
			}
		}
	}

	return dp[(1<<N)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
