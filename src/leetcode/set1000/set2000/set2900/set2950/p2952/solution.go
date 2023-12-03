package p2952

const mod = 1000000007

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func modAdd(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func numberOfSequence(n int, sick []int) int {

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I := make([]int, n+1)
	I[n] = inverse(F[n])

	for i := n - 1; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return modMul(F[n], modMul(I[r], I[n-r]))
	}

	res := 1
	var arr []int
	prev := -1
	var sum int
	for i := 0; i < len(sick); i++ {
		x := sick[i] - prev - 1
		if x > 0 {
			arr = append(arr, x)
			sum += x
			if prev >= 0 {
				res = modMul(res, pow(2, x-1))
			}
		}

		prev = sick[i]
	}

	if prev+1 < n {
		arr = append(arr, n-prev-1)
		sum += n - prev - 1
	}

	for _, x := range arr {
		res = modMul(res, nCr(sum, x))
		sum -= x
	}

	return res
}
