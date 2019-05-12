package p1044

import "sort"

const MOD = 1000000000000007

func longestDupSubstring(S string) string {
	n := len(S)
	check := func(k int) int {
		pow := int64(1)
		hash := make(map[int64]bool)

		if k >= n {
			return -1
		}
		cur := int64(0)
		var i int
		for i < k {
			cur = (cur * 31) + int64(S[i] - 'a')
			cur %= MOD
			pow = (pow * 31) % MOD
			i++
		}

		hash[cur] = true

		for i < n {
			cur = cur * 31 + int64(S[i] - 'a')
			cur = modSub(cur, (int64(S[i-k] - 'a') * pow) % MOD)
			cur %= MOD
			if hash[cur] {
				return i
			}
			hash[cur] = true
			i++
		}
		return -1
	}


	k := sort.Search(n, func(k int) bool {
		return check(k) < 0
	})

	if k <= 1 {
		return ""
	}
	k--
	i := check(k)
	return S[i - k + 1 : i + 1]
}

func modSub(a, b int64) int64 {
	c := a - b
	if c < 0 {
		c += MOD
	}
	return c
}

