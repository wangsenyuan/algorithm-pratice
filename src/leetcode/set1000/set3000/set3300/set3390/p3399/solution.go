package p3399

import (
	"sort"
	"strings"
)

func minLength(s string, numOps int) int {
	// binary search?
	n := len(s)
	check := func(exp int) bool {
		if exp <= 1 {
			return false
		}
		// 能否在限定次数，得到最长的为exp的结果？
		// 但是还有个情况, 000111111 (3,6)
		var cnt int

		for i := 0; i < n; {
			j := i
			for i < n && s[i] == s[j] {
				i++
			}
			tmp := i - j
			if tmp <= exp {
				continue
			}
			if tmp == 2 && i < n && j > 0 {
				return false
			}
			// tmp > exp
			// 假设使用了x次操作, (x + 1) * (exp + 1) - 1 >= tmp
			// (x + 1) = tmp / exp
			x := (tmp + 1 + exp) / (exp + 1)
			cnt += max(x-1, 1)
		}

		return cnt <= numOps
	}

	ans := sort.Search(n, check)
	if ans > 2 {
		return ans
	}

	count := func(x string) int {
		var cnt = 0
		for i := 0; i < n; i++ {
			if s[i] != x[i] {
				cnt++
			}
		}
		return cnt
	}

	if count(strings.Repeat("01", (n+1)/2)) <= numOps ||
		count(strings.Repeat("10", (n+1)/2)) <= numOps {
		return 1
	}

	return ans
}
