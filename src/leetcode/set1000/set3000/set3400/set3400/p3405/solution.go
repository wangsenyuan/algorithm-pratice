package p3405

const mod = 1e9 + 7

func mul(arr ...int) int {
	r := 1
	for _, num := range arr {
		r = r * num % mod
	}
	return r
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func countGoodArrays(n int, m int, k int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = mul(i, f[i-1])
	}
	iv := make([]int, n+1)
	iv[n] = inverse(f[n])
	for i := n - 1; i >= 0; i-- {
		iv[i] = mul(i+1, iv[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(f[n], iv[r], iv[n-r])
	}
	// 共有n-1的相邻对，其中k个是相同，那么剩下n-1-k个就是不相同
	//
	res := nCr(n-1, n-1-k)

	res = mul(res, m, pow(m-1, n-1-k))

	return res
}
