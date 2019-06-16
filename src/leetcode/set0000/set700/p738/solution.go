package p738

func monotoneIncreasingDigits(N int) int {
	if N == 0 {
		return 0
	}
	//at most ten digits
	dp := make([]int, 10)
	i := 0
	for x := N; x > 0; x /= 10 {
		dp[i] = x % 10
		i++
	}

	for j, k := 0, i-1; j < k; j, k = j+1, k-1 {
		dp[j], dp[k] = dp[k], dp[j]
	}

	for j := i - 2; j >= 0; j-- {
		if dp[j] > dp[j+1] {
			if dp[j] == 0 {
				//need to borrow one from previous digits
				a := j
				for a > 0 && dp[a] == 0 {
					dp[a] = 9
					a--
				}
				dp[a]--
				for k := j + 1; k < i && dp[k] != 9; k++ {
					dp[k] = 9
				}
				j = a
			} else {
				for k := j + 1; k < i && dp[k] != 9; k++ {
					dp[k] = 9
				}
				dp[j]--
			}
		}
	}

	var res int

	for j := 0; j < i; j++ {
		res = res*10 + dp[j]
	}

	return res
}
