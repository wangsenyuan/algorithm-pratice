package p2483

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func countPalindromes(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 122)
	}

	cnt := make([]int, 11)

	for i := 0; i < n; i++ {
		if i > 0 {
			for x := 1; x <= 10; x++ {
				for y := 1; y <= 10; y++ {
					dp[i][x*11+y] = dp[i-1][x*11+y]
				}
			}
		}
		y := int(s[i]-'0') + 1
		for x := 1; x <= 10; x++ {
			dp[i][x*11+y] = modAdd(dp[i][x*11+y], cnt[x])
		}

		cnt[y]++
	}

	for i := 0; i <= 10; i++ {
		cnt[i] = 0
	}

	fp := make([]int, 122)

	var res int

	for i := n - 1; i > 0; i-- {
		// if i is mid
		for x := 1; x <= 10; x++ {
			for y := x; y <= 10; y++ {
				tmp := modMul(fp[x*11+y], dp[i-1][y*11+x])
				if x != y {
					tmp = modAdd(tmp, modMul(fp[y*11+x], dp[i-1][x*11+y]))
				}
				res = modAdd(res, tmp)
			}
		}
		y := int(s[i]-'0') + 1
		for x := 1; x <= 10; x++ {
			fp[y*11+x] = modAdd(fp[y*11+x], cnt[x])
		}
		cnt[y]++
	}

	return res
}
