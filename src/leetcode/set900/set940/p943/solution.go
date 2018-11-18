package p943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	n := len(A)
	mp := suffixMap(A)
	total := 1<<uint(n) - 1
	dp := make([][]string, total+1)
	for i := 0; i <= total; i++ {
		dp[i] = make([]string, n)
	}

	var dfs func(mask int, i int) string

	dfs = func(mask int, i int) string {
		unset := total ^ mask
		if len(dp[unset][i]) > 0 {
			return dp[unset][i]
		}

		if mask == total {
			return A[i]
		}

		var best string

		for j := 0; j < n; j++ {
			if mask&(1<<uint(j)) == 0 {
				tmp := dfs(mask|(1<<uint(j)), j)
				k := mp[i][j]
				tmp = A[i] + tmp[k:]
				if len(best) == 0 || len(tmp) < len(best) {
					best = tmp
				}
			}
		}
		dp[unset][i] = best
		return best
	}

	var res string

	for i := 0; i < n; i++ {
		tmp := dfs(0, i)
		if len(res) == 0 || len(tmp) < len(res) {
			res = tmp
		}
	}

	return res
}

func suffixMap(A []string) [][]int {
	n := len(A)

	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mp[i][j] = len(A[i])
				continue
			}
			for k := 1; k <= len(A[j]); k++ {
				if strings.HasSuffix(A[i], A[j][:k]) {
					mp[i][j] = k
				}
			}
		}
	}
	return mp
}

func shortestSuperstring1(A []string) string {
	n := len(A)
	mp := suffixMap(A)

	total := 1<<uint(n) - 1
	var dfs func(mask int, i int, cur string)

	var best string
	dfs = func(mask int, i int, cur string) {
		if mask == total {
			if len(best) == 0 || (len(cur) < len(best)) {
				best = cur
			}
			return
		}
		if len(best) > 0 && len(cur) >= len(best) {
			return
		}

		for j := 0; j < n; j++ {
			if mask&(1<<uint(j)) == 0 {
				k := mp[i][j]
				dfs(mask|(1<<uint(j)), j, cur+A[j][k:])
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(1<<uint(i), i, A[i])
	}

	return best
}
