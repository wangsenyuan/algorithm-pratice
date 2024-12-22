package p3395

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

func mul(arr ...int) int {
	r := 1
	for _, x := range arr {
		r = r * x % mod
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

func subsequencesWithMiddleMode(nums []int) int {
	n := len(nums)

	arr := make([]int, n)
	copy(arr, nums)
	id := make(map[int]int)
	for i := range n {
		if v, ok := id[nums[i]]; !ok {
			arr[i] = len(id)
			id[nums[i]] = arr[i]
		} else {
			arr[i] = v
		}
	}

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	res := nCr(n, 5)

	m := len(id)
	suf := make([]int, m)
	for _, i := range arr {
		suf[i]++
	}

	pref := make([]int, m)
	for p, x := range arr {
		// 如果x在中间，但是它不是一个符合条件的序列
		// x只出现一次的情况
		pref[x]++
		suf[x]--
		p++
		tmp := mul(nCr(p-pref[x], 2), nCr(n-p-suf[x], 2))
		res = sub(res, tmp)
		// x 出现两次，但是它不是众数的情况，有另外一个数y，至少出现了两次
		// 前面出现了一次，或者后面出现了一次
		for y := 0; y < m; y++ {
			if y == x {
				continue
			}
			// 前面后面各一个y，其他的一个数是非x/y的数
			tmp1 := mul(pref[y], suf[y])
			// 前面可以放一个数
			res = sub(res, mul(tmp1, p-(pref[x]+pref[y]), suf[x]))
			res = sub(res, mul(tmp1, (n-p)-(suf[x]+suf[y]), pref[x]-1))

			// 前面两个y, 后面一个y
			tmp2 := mul(nCr(pref[y], 2), nCr(suf[y], 1), suf[x])
			res = sub(res, tmp2)
			tmp3 := mul(nCr(pref[y], 1), nCr(suf[y], 2), pref[x]-1)
			res = sub(res, tmp3)
			// 前面两个y，后面没有y
			tmp4 := mul(nCr(pref[y], 2), suf[x], n-p-suf[x]-suf[y])
			res = sub(res, tmp4)
			// 后面两个y，前面没有y
			tmp5 := mul(pref[x]-1, p-pref[x]-pref[y], nCr(suf[y], 2))
			res = sub(res, tmp5)
		}
	}

	return res
}
