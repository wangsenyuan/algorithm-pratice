package p2513

import "strings"

const MOD = 1000000007

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

func inv(a int) int {
	return pow(a, MOD-2)
}

func countAnagrams(s string) int {
	words := strings.Split(s, " ")

	var max_len int

	for _, w := range words {
		max_len = max(max_len, len(w))
	}

	F := make([]int, max_len+1)
	F[0] = 1
	for i := 1; i <= max_len; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I := make([]int, max_len+1)
	I[max_len] = inv(F[max_len])

	for i := max_len - 1; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || n < 0 {
			return 0
		}

		res := modMul(F[n], I[r])
		res = modMul(res, I[n-r])
		return res
	}

	res := 1

	for _, w := range words {
		freq := wordFreq(w)
		x := len(w)
		for i := 0; i < 26; i++ {
			res = modMul(res, nCr(x, freq[i]))
			x -= freq[i]
		}
	}

	return res
}

func wordFreq(s string) []int {
	res := make([]int, 26)

	for i := 0; i < len(s); i++ {
		res[int(s[i]-'a')]++
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
