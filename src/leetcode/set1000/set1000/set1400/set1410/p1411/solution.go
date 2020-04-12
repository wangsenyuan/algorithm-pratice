package p1411

const MOD = 1000000007

func numOfWays(n int) int {
	a121 := int64(6)
	a123 := int64(6)

	for i := 1; i < n; i++ {
		b121 := 3*a121 + 2*a123
		b123 := 2*a121 + 2*a123
		a121, a123 = b121%MOD, b123%MOD
	}

	return int((a121 + a123) % MOD)
}

func numOfWays1(n int) int {
	// const R = 0
	// const G = 1
	// const Y = 2

	colors := make([][]int, 0, 12)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 3; k++ {
				if k == j {
					continue
				}
				colors = append(colors, []int{i, j, k})
			}
		}
	}

	check := func(i, j int) bool {
		first, second := colors[i], colors[j]
		for k := 0; k < 3; k++ {
			if first[k] == second[k] {
				return false
			}
		}
		return true
	}

	can := make([][]bool, len(colors))
	for i := 0; i < len(colors); i++ {
		can[i] = make([]bool, len(colors))
		for j := 0; j < len(colors); j++ {
			can[i][j] = check(i, j)
		}
	}

	dp := make([]int, len(colors))
	dp2 := make([]int, len(colors))

	for j := 0; j < len(dp); j++ {
		dp[j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < len(dp2); j++ {
			for k := 0; k < len(dp); k++ {
				if can[j][k] {
					dp2[j] += dp[k]
					if dp2[j] >= MOD {
						dp2[j] -= MOD
					}
				}
			}
		}
		for j := 0; j < len(dp2); j++ {
			dp[j] = dp2[j]
			dp2[j] = 0
		}
	}

	var res int

	for j := 0; j < len(dp); j++ {
		res += dp[j]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}
