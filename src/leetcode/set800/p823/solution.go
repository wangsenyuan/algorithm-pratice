package p823

import "sort"

const MOD = 1000000007

func numFactoredBinaryTrees(A []int) int {
	sort.Ints(A)
	n := len(A)
	f := make([]int64, n)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		f[i] = 1
		mp[A[i]] = i
	}

	for i := 1; i < n; i++ {
		a := A[i]
		for j := 0; j < i; j++ {
			b := A[j]
			if a%b == 0 {
				c := a / b
				if k, found := mp[c]; found {
					tmp := (f[j] * f[k]) % MOD
					f[i] = (f[i] + tmp) % MOD
				}
			}
		}
	}

	var ans int64

	for i := 0; i < n; i++ {
		ans = (ans + f[i]) % MOD
	}

	return int(ans)
}
