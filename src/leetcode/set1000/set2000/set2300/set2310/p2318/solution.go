package p2318

const MOD = 1000000007

// 6 * 6 * 6 = 216
const H = 216

func distinctSequences(n int) int {
	dp := make([]int, H)

	if n < 3 {
		return bruteForce(n)
	}
	for x := 0; x < H; x++ {
		a := x / (6 * 6)
		b := x % (6 * 6) / 6
		c := x % 6
		a++
		b++
		c++
		if a == b || b == c || a == c || gcd(a, b) > 1 || gcd(b, c) > 1 {
			continue
		}
		dp[x]++
	}

	fp := make([]int, H)

	for i := 3; i < n; i++ {
		for x := 0; x < H; x++ {
			fp[x] = 0
		}
		for x := 0; x < H; x++ {
			if dp[x] == 0 {
				continue
			}
			// a := x / (6 * 6)
			b := x % (6 * 6) / 6
			c := x % 6
			for d := 0; d < 6; d++ {
				if d == b || d == c || gcd(c+1, d+1) > 1 {
					continue
				}
				nx := x%(36)*6 + d
				fp[nx] = modAdd(fp[nx], dp[x])
			}
		}
		copy(dp, fp)
	}

	var res int
	for x := 0; x < H; x++ {
		res = modAdd(res, dp[x])
	}

	return res
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func bruteForce(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 6)
	}
	for j := 0; j < 6; j++ {
		dp[0][j]++
	}
	for i := 1; i < n; i++ {
		for k := 0; k < 6; k++ {
			for j := 0; j < 6; j++ {
				if j != k && gcd(j+1, k+1) == 1 {
					dp[i][k] += dp[i-1][j]
				}
			}
		}
	}
	var res int

	for j := 0; j < 6; j++ {
		res += dp[n-1][j]
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
