package p891_

import "sort"

const MOD = 1e9 + 7

func sumSubseqWidths(A []int) int {
	n := len(A)
	sort.Ints(A)
	var ans int64
	for i := 0; i < n; i++ {
		a := pow(2, i)
		b := pow(2, n-i-1)
		c := a - b
		if c < 0 {
			c += MOD
		}
		ans += (c * int64(A[i])) % MOD
		if ans >= MOD {
			ans -= MOD
		}
	}

	return int(ans)
}

func pow(a, b int) int64 {
	x := int64(a)
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = (res * x) % MOD
		}
		x = (x * x) % MOD
		b >>= 1
	}
	return res
}
