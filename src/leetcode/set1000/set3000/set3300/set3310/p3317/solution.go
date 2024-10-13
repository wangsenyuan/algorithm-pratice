package p3317

const mod = 1e9 + 7

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
	return a * b % mod
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

func numberOfWays(n int, x int, y int) int {
	S := make([][]int, 1001)
	for i := 0; i <= 1000; i++ {
		S[i] = make([]int, 1001)
	}
	S[0][0] = 1
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			S[i][j] = add(S[i-1][j-1], mul(j, S[i-1][j]))
		}
	}

	var res int
	perm := 1
	pw := 1

	for i := 1; i <= min(n, x); i++ {
		perm = mul(perm, x+1-i)
		pw = mul(pw, y)
		res = add(res, mul(pw, mul(perm, S[n][i])))
	}

	return res
}
