package coins

const MOD = 1000000007

func waysToChange(n int) int {
	if n == 0 {
		return 1
	}
	ps := []int{25, 10, 5, 1}
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < len(ps); i++ {
		x := ps[i]
		for j := x; j <= n; j++ {
			modAdd(&dp[j], dp[j-x])
		}
	}

	return dp[n]
}

func waysToChange2(n int) int {
	if n == 0 {
		return 1
	}
	ps := []int{0, 1, 5, 10, 25}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, len(ps))
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < len(ps); j++ {
			//dp[i][j]
			x := ps[j]
			if i < x {
				break
			}
			if i == x {
				modAdd(&dp[i][j], 1)
			} else {
				modAdd(&dp[i][j], dp[i-x][j])
			}
		}
		for j := 1; j < len(ps); j++ {
			modAdd(&dp[i][j], dp[i][j-1])
		}
	}

	var res int = dp[n][4]

	// for j := 1; j < len(ps); j++ {
	// 	modAdd(&res, dp[n][j])
	// }

	return res
}

func waysToChange1(n int) int {
	if n == 0 {
		return 1
	}
	ps := []int{0, 1, 5, 10, 25}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, len(ps))
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < len(ps); j++ {
			//dp[i][j]
			x := ps[j]
			if i < x {
				break
			}
			for k := 0; k <= j; k++ {
				modAdd(&dp[i][j], dp[i-x][k])
			}
		}
	}

	var res int

	for j := 1; j < len(ps); j++ {
		modAdd(&res, dp[n][j])
	}

	return res
}

func modAdd(a *int, b int) {
	*a += b

	if *a >= MOD {
		*a -= MOD
	}
}
