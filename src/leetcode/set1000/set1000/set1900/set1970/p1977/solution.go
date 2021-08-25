package p1977

const MOD = 1000000007

func numberOfCombinations(num string) int {
	if num[0] == '0' {
		return 0
	}

	n := len(num)
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			// fp[i][j] = first different character away from i & j
			if num[i] == num[j] {
				if j+1 < n {
					fp[i][j] = fp[i+1][j+1] + 1
				} else {
					fp[i][j] = 0
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			continue
		}
		for j := n - 1; j >= i; j-- {
			// dp[i][j] += dp[i][j+1]
			// k - (j + 1) + 1 = j - i + 1
			if j == n-1 {
				dp[i][j] = 1
			} else {
				k := j + 1 + j - i
				if k < n {
					x := fp[i][j+1]
					if i+x <= j && num[i+x] > num[j+1+x] {
						k++
					}
					if k < n {
						modAdd(&dp[i][j], dp[j+1][k])
					}
				}
			}

			if j-1 >= i {
				modAdd(&dp[i][j-1], dp[i][j])
			}
		}
	}

	return dp[0][0]
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func numberOfCombinations1(num string) int {
	if num[0] == '0' {
		return 0
	}
	n := len(num)
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			// fp[i][j] = first different character away from i & j
			if num[i] == num[j] {
				if j+1 < n {
					fp[i][j] = fp[i+1][j+1] + 1
				} else {
					fp[i][j] = 0
				}
			}
		}
	}
	var fn func(i, j int) int
	fn = func(i, j int) int {
		if num[i] == '0' {
			return 0
		}
		if j == n-1 {
			return 1
		}

		if dp[i][j] >= 0 {
			return dp[i][j]
		}

		dp[i][j] = 0

		if j+1 < n {
			modAdd(&dp[i][j], fn(i, j+1))
		}

		//find comb numbers that starts from num[i:j+1]
		// k - (j + 1) + 1 = j - i + 1
		k := j + 1 + j - i

		if k >= n {
			// num[j+1:] < num[i:j+1]
			return dp[i][j]
		}

		l := fp[i][j+1]

		if i+l <= j && num[i+l] > num[j+1+l] {
			// num[i:j+1] > num[j+1:k+1]
			k++
		}
		if k >= n {
			return dp[i][j]
		}

		modAdd(&dp[i][j], fn(j+1, k))

		return dp[i][j]
	}

	return fn(0, 0)
}
