package p3428

import "sort"

func minMaxSums(nums []int, k int) int {
	n := len(nums)

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
		return mul(F[n], I[r], I[n-r])
	}

	var res int

	sort.Ints(nums)

	for i, x := range nums {
		var sum int
		for j := 0; j < min(k, i+1); j++ {
			sum = add(sum, nCr(i, j))
		}

		res = add(res, mul(x, sum))
		res = add(res, mul(nums[n-1-i], sum))
	}

	return res
}

const mod = 1e9 + 7

func add(a, b int) int {
	if a < 0 {
		a += mod
	}
	if b < 0 {
		b += mod
	}
	a += b
	if a >= mod {
		a -= mod
	}
	return a
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
