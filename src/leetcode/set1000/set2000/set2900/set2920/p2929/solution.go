package p2929

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func stringCount(n int) int {
	if n < 4 {
		return 0
	}

	ans := pow(26, n)
	ans = sub(ans, mul(75+n, pow(25, n-1)))

	ans = add(ans, mul(72+2*n, pow(24, n-1)))
	// 不包含l, 不包含t， 不包含e，或者只有一个e

	ans = sub(ans, mul(23+n, pow(23, n-1)))

	return ans
}
