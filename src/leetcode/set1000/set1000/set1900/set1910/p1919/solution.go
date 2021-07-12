package p1919

const MOD = 1000000007

func colorTheGrid(m int, n int) int {
	// m <= 5 && n <= 1000
	M := pow(3, m)

	arr := make([]int, m)
	rra := make([]int, m)

	check := func(a, b int) bool {
		deNum(a, m, arr)
		deNum(b, m, rra)

		for i := 0; i < m; i++ {
			if arr[i] == rra[i] {
				return false
			}
			if i > 0 && (arr[i] == arr[i-1] || rra[i] == rra[i-1]) {
				return false
			}
		}
		return true
	}

	dp := make([]int, M)
	fp := make([]int, M)
	good := make([]int, 0, M)

	for i := 0; i < M; i++ {
		deNum(i, m, arr)
		ok := true
		for j := 1; j < m; j++ {
			if arr[j] == arr[j-1] {
				ok = false
				break
			}
		}
		if ok {
			dp[i] = 1
			good = append(good, i)
		}
	}

	can := make([][]int, M)
	for i := 0; i < M; i++ {
		can[i] = make([]int, 0, M)
		for _, j := range good {
			if check(i, j) {
				can[i] = append(can[i], j)
			}
		}
	}

	for i := 1; i < n; i++ {
		for _, state := range good {
			fp[state] = 0
			for _, j := range can[state] {
				fp[state] += dp[j]
				if fp[state] >= MOD {
					fp[state] -= MOD
				}
			}
		}
		copy(dp, fp)
	}
	var res int
	for i := 0; i < M; i++ {
		res += dp[i]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func deNum(num int, m int, arr []int) {
	for i := m - 1; i >= 0; i-- {
		arr[i] = num % 3
		num /= 3
	}
}

func pow(a, b int) int {
	A := int64(a)
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(res)
}
